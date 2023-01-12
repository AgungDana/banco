package repo

import (
	"banco/domain/user"
	"context"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) user.Query {
	return &query{db: db}
}

// GetUserByEmail implements user.Query
func (q *query) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	person := new(user.User)
	err := q.db.Where("email = ?", email).Find(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}

// GetUserById implements user.Query
func (q *query) FindUserByID(ctx context.Context, id uint) (*user.User, error) {
	person := new(user.User)
	err := q.db.Where("id = ?", id).Find(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}
