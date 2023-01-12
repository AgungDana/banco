package repo

import (
	"banco/common/infra/orm"
	"banco/domain/produk"
	"context"
	"log"
	"sync"

	"gorm.io/gorm"
)

type mutation struct {
	txn   *gorm.DB
	query produk.Query
	mu    sync.Mutex
}

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
func (m *mutation) reset(ctx context.Context) {
	m.txn = nil
}

func newMutation(ctx context.Context, db *gorm.DB) produk.Mutation {
	txn := orm.BeginTx(ctx, db)
	return &mutation{txn: txn, query: newQuery(txn)}
}

// AddAmoutProduct implements produk.Mutation
func (m *mutation) AddAmout(ctx context.Context, amount produk.Amount) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Create(&amount).Error
	if err != nil {
		return nil, err
	}
	return &amount.Id, nil
}

// AddSatuan implements produk.Mutation
func (m *mutation) AddSatuan(ctx context.Context, satuan produk.Satuan) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Create(&satuan).Error
	if err != nil {
		return nil, err
	}
	return &satuan.Id, nil
}

// CreateProduct implements produk.Mutation
func (m *mutation) CreateProduct(ctx context.Context, p produk.Product) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return &p.Id, nil
}

// CreateProductType implements produk.Mutation
func (m *mutation) CreateProductType(ctx context.Context, p produk.ProductType) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return &p.Id, nil
}

// DeleteAmoutProduct implements produk.Mutation
func (m *mutation) DeleteAmout(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Delete(produk.Amount{}, id).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteProduct implements produk.Mutation
func (m *mutation) DeleteProduct(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Delete(produk.Product{}, id).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteProductType implements produk.Mutation
func (m *mutation) DeleteProductType(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Delete(produk.ProductType{}, id).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// DeleteSatuan implements produk.Mutation
func (m *mutation) DeleteSatuan(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Delete(produk.Satuan{}, id).Error
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UpdateAmoutProduct implements produk.Mutation
func (m *mutation) UpdateAmout(ctx context.Context, amount produk.Amount) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Where("id = ?", amount.Id).Updates(&amount).Error
	if err != nil {
		return nil, err
	}
	return &amount.Id, nil
}

// UpdateProduct implements produk.Mutation
func (m *mutation) UpdateProduct(ctx context.Context, p produk.Product) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Where("id = ?", p.Id).Updates(&p).Error
	if err != nil {
		m.txn.Rollback()
		return nil, err
	}
	return &p.Id, nil
}

// UpdateProductType implements produk.Mutation
func (m *mutation) UpdateProductType(ctx context.Context, p produk.ProductType) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Where("id = ?", p.Id).Updates(&p).Error
	if err != nil {
		m.txn.Rollback()
		return nil, err
	}
	return &p.Id, nil
}

// UpdateSatuan implements produk.Mutation
func (m *mutation) UpdateSatuan(ctx context.Context, satuan produk.Satuan) (*uint, error) {
	// panic("unimplemented")
	err := m.txn.Where("id = ?", satuan.Id).Updates(&satuan).Error
	if err != nil {
		m.txn.Rollback()
		return nil, err
	}
	return &satuan.Id, nil
}
