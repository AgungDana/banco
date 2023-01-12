package manager

import (
	"banco/common/config"
	"banco/common/ctxutil"
	inveRepo "banco/domain/inventori/repo"
	"banco/domain/transaksi"
	"banco/domain/transaksi/repo"
	"context"
)

type manager struct {
	repo transaksi.Repo
}

func NewManager(conf config.Config) transaksi.Manager {
	inveRepo, err := inveRepo.NewRepoInventory(conf)
	if err != nil {
		panic(err)
	}
	transaksiRepo, err := repo.NewRepo(conf, inveRepo)
	if err != nil {
		panic(err)
	}
	return &manager{repo: transaksiRepo}
}

func (m *manager) FindTrasaksiByID(ctx context.Context, id uint) (*transaksi.Transaction, error) {
	repo := m.repo.NewQuery()
	data, err := repo.FindTrasaksiByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *manager) FindTrasaksiByProductID(ctx context.Context, id uint) (*transaksi.Transaction, error) {
	repo := m.repo.NewQuery()
	data, err := repo.FindTrasaksiByProductID(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *manager) FindAllTrasaksi(ctx context.Context) ([]*transaksi.Transaction, error) {
	repo := m.repo.NewQuery()
	data, err := repo.FindAllTrasaksi(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (m *manager) CreateTransaksi(ctx context.Context, txn transaksi.TransactionReq) (*uint, error) {
	repo := m.repo.NewMutation(ctx)

	userId, _ := ctxutil.GetUserIdFromCtx(ctx)

	newTxn := transaksi.ToTransaction(txn, *userId)
	data, err := repo.CreateTransaksi(ctx, newTxn)
	if err != nil {
		repo.Cancel(ctx)
		return nil, err
	}
	if err = repo.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}
