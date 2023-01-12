package repo

import (
	"banco/domain/transaksi"
	"context"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) transaksi.Query {
	return &query{db: db}
}

func (q *query) FindTrasaksiByID(ctx context.Context, id uint) (*transaksi.Transaction, error) {
	data := transaksi.Transaction{}
	err := q.db.Where("id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (q *query) FindTrasaksiByProductID(ctx context.Context, id uint) (*transaksi.Transaction, error) {
	data := transaksi.Transaction{}
	err := q.db.Where("product_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (q *query) FindAllTrasaksi(ctx context.Context) ([]*transaksi.Transaction, error) {
	data := []*transaksi.Transaction{}
	err := q.db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
