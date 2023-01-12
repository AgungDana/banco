package repo

import (
	"banco/common/infra/orm"
	"banco/common/werror"
	"banco/domain/inventori"
	"banco/domain/transaksi"
	"context"
	"errors"
	"sync"

	"gorm.io/gorm"
)

type mutation struct {
	txn        *gorm.DB
	mu         sync.Mutex
	query      transaksi.Query
	invenMut   inventori.Mutation
	invenQeury inventori.Query
}

func newMutation(ctx context.Context, db *gorm.DB, inven inventori.Repo) transaksi.Mutation {
	txn := orm.BeginTx(ctx, db)
	inventoriMutation := inven.NewMutation(ctx)
	inventoriQuery := inven.NewQuery()
	return &mutation{txn: txn, query: newQuery(txn), invenMut: inventoriMutation, invenQeury: inventoriQuery}
}

func (m *mutation) Commit(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer m.reset()

	if err := orm.ComitTx(ctx); err != nil {
		return errors.New("invalid server error")
	}

	return m.invenMut.Commit(ctx)
}

func (m *mutation) Cancel(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	defer m.reset()

	if err := orm.RoolBackTxn(ctx); err != nil {
		return err
	}

	return nil
}

func (m *mutation) reset() {
	m.txn = nil
}

func (m *mutation) CreateTransaksi(ctx context.Context, txn transaksi.Transaction) (*uint, error) {

	err := m.txn.Create(&txn).Error
	if err != nil {
		return nil, err
	}

	errs := werror.New()

	for _, v := range txn.TransactionDetail {
		data, err := m.invenQeury.FindProductFromInventoryByIdProduct(ctx, v.ProductId)
		if err != nil {
			errs.AddError(errors.New("product Tidak ada :" + v.ProductName))
			break
		}
		if data.Stock-v.TotalItem <= 0 {
			errs.AddError(errors.New("stock di invetory tidak mencukupi :" + v.ProductName))
			break
		}

		data.Stock = data.Stock - v.TotalItem
		_, err = m.invenMut.UpdateInventory(ctx, *data)
		if err != nil {
			errs.AddError(errors.New("failed update inventory stock :" + v.ProductName))
			return nil, err
		}
	}

	return &txn.Id, nil
}
