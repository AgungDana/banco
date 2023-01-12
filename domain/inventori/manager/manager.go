package manager

import (
	"banco/common/config"
	"banco/common/ctxutil"
	"banco/common/infra/orm"
	"banco/domain/inventori"
	"banco/domain/inventori/repo"

	"context"
)

type manager struct {
	repo inventori.Repo
}

func NewManager(conf config.Config) inventori.Manager {
	repo, _ := repo.NewRepoInventory(conf)
	return &manager{
		repo: repo,
	}
}

// FindAll implements inventori.Manager
func (m *manager) FindAll(ctx context.Context) ([]*inventori.Inventory, error) {
	repo := m.repo.NewQuery()
	data, err := repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return data, err
}

// FindProductFromInventoryByIdInventory implements inventori.Manager
func (m *manager) FindProductFromInventoryByIdInventory(ctx context.Context, id uint) (*inventori.Inventory, error) {
	repo := m.repo.NewQuery()

	data, err := repo.FindProductFromInventoryByIdInventory(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// FindProductFromInventoryByIdProduct implements inventori.Manager
func (m *manager) FindProductFromInventoryByIdProduct(ctx context.Context, id uint) (*inventori.Inventory, error) {
	repo := m.repo.NewQuery()

	data, err := repo.FindProductFromInventoryByIdProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// AddProductToInventory implements inventori.Manager
func (m *manager) AddProductToInventory(ctx context.Context, req inventori.CreateRequest) (*uint, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProductFromInventoryByIdProduct(ctx, req.ProductId)
	if err != nil {
		return nil, err
	}

	id, err := ctxutil.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	mut := m.repo.NewMutation(ctx)
	if data != nil {
		stock := data.Stock + req.Stock
		incoming := data.IncomingItem + req.IncomingItem

		add := inventori.Inventory{
			Model:        orm.Model{Id: data.Id, UpdatedBy: *id},
			ProductId:    req.ProductId,
			IncomingItem: incoming,
			Stock:        stock,
			Balance:      0,
			SatuanId:     req.SatuanId,
		}
		data, err := mut.UpdateInventory(ctx, add)
		if err != nil {
			return nil, err
		}
		return data, nil
	} else {
		add := inventori.Inventory{
			Model: orm.Model{
				Id:        data.Id,
				CraetedBy: *id,
			},
			ProductId:    req.ProductId,
			IncomingItem: req.IncomingItem,
			Stock:        req.Stock,
			SatuanId:     req.SatuanId,
		}
		data, err := mut.AddProductToInventory(ctx, add)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}

// IncreaseProductInInventory implements inventori.Manager
func (m *manager) IncreaseProductInInventory(ctx context.Context, inve inventori.InventoryIncreaseRequest) (*uint, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProductFromInventoryByIdInventory(ctx, inve.Id)
	if err != nil {
		return nil, err
	}

	id, err := ctxutil.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req := inventori.IncreaseReqToInventory(inve)

	req.UpdatedBy = *id
	req.Stock = req.Stock + data.Stock
	req.IncomingItem = req.IncomingItem + data.IncomingItem

	mut := m.repo.NewMutation(ctx)
	id, err = mut.UpdateInventory(ctx, req)
	if err != nil {
		mut.Cancel(ctx)
		return nil, err
	}

	if err = mut.Commit(ctx); err != nil {
		return nil, err
	}
	return id, nil
}

// DecreaseProductInInventory implements inventori.Manager
func (m *manager) DecreaseProductInInventory(ctx context.Context, inve inventori.InventoryDecreaseRequest) (*uint, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProductFromInventoryByIdInventory(ctx, inve.Id)
	if err != nil {
		return nil, err
	}

	id, err := ctxutil.GetUserIdFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	req := inventori.DecreaseReqToInventory(inve)

	req.UpdatedBy = *id
	req.Stock = data.Stock - req.Stock
	req.OutcomingItem = req.OutcomingItem + data.OutcomingItem

	mut := m.repo.NewMutation(ctx)

	id, err = mut.UpdateInventory(ctx, req)
	if err != nil {
		mut.Cancel(ctx)
		return nil, err
	}

	if err = mut.Commit(ctx); err != nil {
		return nil, err
	}
	return id, nil
}

// DeleteProductFromInvetory implements inventori.Manager
func (m *manager) DeleteProductFromInvetory(ctx context.Context, id uint) (*uint, error) {
	repo := m.repo.NewMutation(ctx)

	data, err := repo.DeleteProductFromInventory(ctx, id)
	if err != nil {
		repo.Cancel(ctx)
		return nil, err
	}

	if err = repo.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}
