package transaksi

import "context"

type Manager interface {
	FindTrasaksiByID(ctx context.Context, id uint) (*Transaction, error)
	FindTrasaksiByProductID(ctx context.Context, id uint) (*Transaction, error)
	FindAllTrasaksi(ctx context.Context) ([]*Transaction, error)
	CreateTransaksi(ctx context.Context, txn TransactionReq) (*uint, error)
}

type Repo interface {
	NewMutation(ctx context.Context) Mutation
	NewQuery() Query
}

type Mutation interface {
	CreateTransaksi(ctx context.Context, txn Transaction) (*uint, error)

	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}

type Query interface {
	FindTrasaksiByID(ctx context.Context, id uint) (*Transaction, error)
	FindTrasaksiByProductID(ctx context.Context, id uint) (*Transaction, error)
	FindAllTrasaksi(ctx context.Context) ([]*Transaction, error)
}
