package repo

import (
	"banco/domain/produk"
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) produk.Query { //abstrak
	return &query{db: db}
}

// FindAmount implements produk.Query
func (q *query) FindAmount(ctx context.Context, id uint) (*produk.Amount, error) {
	// panic("unimplemented")
	p := produk.Amount{}
	err := q.db.Where("id=?", id).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// FindAmounts implements produk.Query
func (q *query) FindAmounts(ctx context.Context) ([]*produk.Amount, error) {
	// panic("unimplemented")
	p := []*produk.Amount{}
	// err := q.db.Where("time_start >= ? AND time_end <=  ?", time.Now(), time.Now()).Find(&p).Error
	err := q.db.Find(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	if len(p) == 0 {
		return nil, errors.New("not found")
	}
	return p, nil
}

// FindProduct implements produk.Query
func (q *query) FindProduct(ctx context.Context, id uint) (*produk.Product, error) {
	// panic("unimplemented")
	p := produk.Product{}
	err := q.db.Where("id=?", id).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, err
}

// FindProductType implements produk.Query
func (q *query) FindProductType(ctx context.Context, id uint) (*produk.ProductType, error) {
	// panic("unimplemented")
	p := produk.ProductType{}
	err := q.db.Where("id=?", id).Preload("Product").Find(&p).Error
	fmt.Printf("err: %v\n", err)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// FindProductTypes implements produk.Query
func (q *query) FindProductTypes(ctx context.Context) ([]*produk.ProductType, error) {
	// panic("unimplemented")
	p := []*produk.ProductType{}
	// err := q.db.Where("time_start >= ? AND time_end <= ?", time.Now(), time.Now()).Find(&p).Error
	err := q.db.Find(&p).Error
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, errors.New("not found")
	}
	return p, nil
}

// FindProducts implements produk.Query
func (q *query) FindProducts(ctx context.Context) ([]*produk.Product, error) {
	// panic("unimplemented")
	p := []*produk.Product{}
	// err := q.db.Where("time_start >= ? AND time_end <= ?", time.Now(), time.Now()).Find(&p).Error
	err := q.db.Find(&p).Error
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, errors.New("not found")
	}
	return p, nil
}

// FindSatuan implements produk.Query
func (q *query) FindSatuan(ctx context.Context, id uint) (*produk.Satuan, error) {
	// panic("unimplemented")
	p := produk.Satuan{}
	err := q.db.Where("id=?", id).Find(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// FindSatuans implements produk.Query
func (q *query) FindSatuans(ctx context.Context) ([]*produk.Satuan, error) {
	// panic("unimplemented")
	p := []*produk.Satuan{}
	err := q.db.Find(&p).Error
	if err != nil {
		return nil, err
	}
	if len(p) == 0 {
		return nil, errors.New("not found")
	}
	return p, nil
}
