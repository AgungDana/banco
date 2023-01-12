package produk

import "context"

type Manager interface {
	GetProduct(ctx context.Context, id uint) (*Product, error)
	GetProducts(ctx context.Context) ([]*Product, error)
	CreateProduct(ctx context.Context, p Product) (*Product, error)
	UpdateProduct(ctx context.Context, p Product) (*Product, error)
	DeleteProduct(ctx context.Context, id uint) (*uint, error)

	GetProductType(ctx context.Context, id uint) (*ProductType, error)
	GetProductTypes(ctx context.Context) ([]*ProductType, error)
	CreateProductType(ctx context.Context, p ProductType) (*ProductType, error)
	UpdateProductType(ctx context.Context, p ProductType) (*ProductType, error)
	DeleteProductType(ctx context.Context, id uint) (*uint, error)

	GetAmout(ctx context.Context, id uint) (*Amount, error)
	GetAmouts(ctx context.Context) ([]*Amount, error)
	CreateAmout(ctx context.Context, a Amount) (*Amount, error)
	UpdateAmout(ctx context.Context, a Amount) (*Amount, error)
	DeleteAmout(ctx context.Context, id uint) (*uint, error)

	GetSatuan(ctx context.Context, id uint) (*Satuan, error)
	GetSatuans(ctx context.Context) ([]*Satuan, error)
	CreateSatuan(ctx context.Context, s Satuan) (*Satuan, error)
	UpdateSatuan(ctx context.Context, s Satuan) (*Satuan, error)
	DeleteSatuan(ctx context.Context, id uint) (*uint, error)
}

type Repo interface {
	NewMutation(ctx context.Context) Mutation
	NewQuery() Query
}

type Mutation interface {
	CreateProduct(ctx context.Context, p Product) (*uint, error)
	UpdateProduct(ctx context.Context, p Product) (*uint, error)
	DeleteProduct(ctx context.Context, id uint) (*uint, error)

	CreateProductType(ctx context.Context, p ProductType) (*uint, error)
	UpdateProductType(ctx context.Context, p ProductType) (*uint, error)
	DeleteProductType(ctx context.Context, id uint) (*uint, error)

	AddAmout(ctx context.Context, amout Amount) (*uint, error)
	UpdateAmout(ctx context.Context, amout Amount) (*uint, error)
	DeleteAmout(ctx context.Context, id uint) (*uint, error)

	AddSatuan(ctx context.Context, satuan Satuan) (*uint, error)
	UpdateSatuan(ctx context.Context, satuan Satuan) (*uint, error)
	DeleteSatuan(ctx context.Context, id uint) (*uint, error)

	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}

type Query interface {
	FindProduct(ctx context.Context, id uint) (*Product, error)
	FindProducts(ctx context.Context) ([]*Product, error)

	FindProductType(ctx context.Context, id uint) (*ProductType, error)
	FindProductTypes(ctx context.Context) ([]*ProductType, error)

	FindAmount(ctx context.Context, id uint) (*Amount, error)
	FindAmounts(ctx context.Context) ([]*Amount, error)

	FindSatuan(ctx context.Context, id uint) (*Satuan, error)
	FindSatuans(ctx context.Context) ([]*Satuan, error)
}
