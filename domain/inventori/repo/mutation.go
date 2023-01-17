package repo

import (
	"banco/common/infra/orm"
	"banco/domain/inventori"
	"context"
	"fmt"
	"sync"

	"gorm.io/gorm"
)

type mutation struct {
	txn   *gorm.DB
	query inventori.Query
	mu    sync.Mutex
}

func newMutation(ctx context.Context, db *gorm.DB) inventori.Mutation {
	txn := orm.BeginTx(ctx, db)
	return &mutation{txn: txn, query: newQuery(txn)}
}

func (m *mutation) Cancel(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := orm.RoolBackTxn(ctx); err != nil {
		return err
	}
	return nil
}

func (m *mutation) Commit(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if err := orm.ComitTx(ctx); err != nil {
		return err
	}
	return nil
}

// AddProductToInventory implements inventori.Mutation
func (m *mutation) AddProductToInventory(ctx context.Context, inve inventori.Inventory) (*uint, error) {
	fmt.Printf("inve: %v\n", inve)
	err := m.txn.Model(inventori.Inventory{}).Create(&inve).Error
	if err != nil {
		return nil, err
	}
	return &inve.Id, nil
}

// UpdateInventory implements inventori.Mutation
func (m *mutation) UpdateInventory(ctx context.Context, inve inventori.Inventory) (*uint, error) {
	err := m.txn.Model(inventori.Inventory{}).Where("id = ?", inve.Id).Updates(&inve).Error
	if err != nil {
		return nil, err
	}
	return &inve.Id, nil
}

// DeleteProductFromInventory implements inventori.Mutation
func (m *mutation) DeleteProductFromInventory(ctx context.Context, id uint) (*uint, error) {
	err := m.txn.Delete(&inventori.Inventory{}, id).Error
	if err != nil {
		return nil, err
	}
	return nil, err
}
