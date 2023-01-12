package inventori

import "context"

type Manager interface {
	FindProductFromInventoryByIdInventory(ctx context.Context, id uint) (*Inventory, error)
	FindProductFromInventoryByIdProduct(ctx context.Context, id uint) (*Inventory, error)
	FindAll(ctx context.Context) ([]*Inventory, error)
	AddProductToInventory(ctx context.Context, inventori CreateRequest) (*uint, error)
	IncreaseProductInInventory(ctx context.Context, inve InventoryIncreaseRequest) (*uint, error)
	DecreaseProductInInventory(ctx context.Context, inve InventoryDecreaseRequest) (*uint, error)
	DeleteProductFromInvetory(ctx context.Context, id uint) (*uint, error)
}

type Repo interface {
	NewMutation(ctx context.Context) Mutation
	NewQuery() Query
}

type Mutation interface {
	AddProductToInventory(ctx context.Context, inventori Inventory) (*uint, error)
	UpdateInventory(ctx context.Context, inventori Inventory) (*uint, error)
	DeleteProductFromInventory(ctx context.Context, id uint) (*uint, error)
	Commit(ctx context.Context) error
	Cancel(ctx context.Context) error
}

type Query interface {
	FindProductFromInventoryByIdInventory(ctx context.Context, id uint) (*Inventory, error)
	FindProductFromInventoryByIdProduct(ctx context.Context, id uint) (*Inventory, error)
	FindAll(ctx context.Context) ([]*Inventory, error)
}
