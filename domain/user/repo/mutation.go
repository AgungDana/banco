package repo

import (
	"banco/common/infra/orm"
	"banco/domain/user"
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"gorm.io/gorm"
)

type mutation struct {
	txn   *gorm.DB
	query user.Query
	mu    sync.Mutex
}

func newMutation(ctx context.Context, db *gorm.DB) user.Mutation {
	var (
		txn = orm.BeginTx(ctx, db)
	)
	return &mutation{txn: txn, query: newQuery(txn)}
}

// Cancel implements user.Mutation
func (m *mutation) Cancel(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer m.reset(ctx)

	if err := orm.RoolBackTxn(ctx); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Commit implements user.Mutation
func (m *mutation) Commit(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer m.reset(ctx)
	if err := orm.ComitTx(ctx); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// CreateUser implements user.Mutation
func (m *mutation) CreateUser(ctx context.Context, userReq user.User) (*uint, error) {
	person, _ := m.query.FindUserByEmail(ctx, userReq.Email)
	if person.Email == userReq.Email {
		fmt.Println(person)
		return nil, errors.New("email already use")
	}

	err := m.txn.Create(&userReq).Error
	if err != nil {
		return nil, err
	}
	return &userReq.Id, nil
}

// DeletUserByID implements user.Mutation
func (m *mutation) DeletUserByID(ctx context.Context, id uint) error {
	err := m.txn.Delete(user.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateUserByEmail implements user.Mutation
func (m *mutation) UpdateUserByEmail(ctx context.Context, update user.User) (*uint, error) {
	err := m.txn.Where("email = ?", update.Email).Updates(&update).Error
	if err != nil {
		return nil, err
	}
	return &update.Id, nil
}

// UpdateUserByID implements user.Mutation
func (m *mutation) UpdateUserByID(ctx context.Context, update user.User) (*uint, error) {
	err := m.txn.Where("id = ?", update.Id).Updates(&update).Error
	if err != nil {
		m.txn.Rollback()
		return nil, err
	}
	return &update.Id, nil
}

func (m *mutation) reset(ctx context.Context) {
	m.txn = nil
}
