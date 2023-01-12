package repo

import (
	"banco/domain/inventori"
	"context"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) inventori.Query {
	return &query{db: db}
}

// FindAll implements inventori.Query
func (q *query) FindAll(ctx context.Context) ([]*inventori.Inventory, error) {
	data := []*inventori.Inventory{}
	err := q.db.Find(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

// FindProductFromInventoryByIdInventory implements inventori.Query
func (q *query) FindProductFromInventoryByIdInventory(ctx context.Context, id uint) (*inventori.Inventory, error) {
	data := inventori.Inventory{}
	err := q.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

// FindProductFromInventoryByIdProduct implements inventori.Query
func (q *query) FindProductFromInventoryByIdProduct(ctx context.Context, id uint) (*inventori.Inventory, error) {
	data := inventori.Inventory{}
	err := q.db.Where("product_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
