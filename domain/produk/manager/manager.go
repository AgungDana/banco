package manager

import (
	"banco/common/config"
	"banco/common/ctxutil"
	"banco/domain/produk"
	"banco/domain/produk/repo"
	"context"
	"log"
)

type manager struct {
	repo produk.Repo
}

func NewManager(conf config.Config) produk.Manager {
	// repo, _ := produk.Repo
	repo, err := repo.NewRepoProduct(conf)
	if err != nil {
		// return nil, err
		panic(err)
	}
	return &manager{repo: repo}
}

// CreateAmout implements produk.Manager
func (m *manager) CreateAmout(ctx context.Context, a produk.Amount) (*produk.Amount, error) {
	// panic("unimplemented")
	// mut, err := produk.Mutation.AddAmoutProduct(ctx, a)
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.AddAmout(ctx, a)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindAmount(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateProduct implements produk.Manager
func (m *manager) CreateProduct(ctx context.Context, p produk.CreateProductRequest) (*produk.Product, error) {
	// panic("unimplemented")
	userId, _ := ctxutil.GetUserIdFromCtx(ctx)
	mutation := m.repo.NewMutation(ctx)
	product := produk.NewProduct(p, *userId)
	query := m.repo.NewQuery()
	id, err := mutation.CreateProduct(ctx, product)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindProduct(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// CreateProductType implements produk.Manager
func (m *manager) CreateProductType(ctx context.Context, p produk.ProductType) (*produk.ProductType, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.CreateProductType(ctx, p)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindProductType(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateSatuan implements produk.Manager
func (m *manager) CreateSatuan(ctx context.Context, s produk.Satuan) (*produk.Satuan, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.AddSatuan(ctx, s)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindSatuan(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// DeleteAmout implements produk.Manager
func (m *manager) DeleteAmout(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	data, err := mutation.DeleteAmout(ctx, id)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil

}

// DeleteProduct implements produk.Manager
func (m *manager) DeleteProduct(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	data, err := mutation.DeleteProduct(ctx, id)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil

}

// DeleteProductType implements produk.Manager
func (m *manager) DeleteProductType(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	data, err := mutation.DeleteProductType(ctx, id)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil

}

// DeleteSatuan implements produk.Manager
func (m *manager) DeleteSatuan(ctx context.Context, id uint) (*uint, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	data, err := mutation.DeleteSatuan(ctx, id)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	return data, nil
}

// GetAmout implements produk.Manager
func (m *manager) GetAmout(ctx context.Context, id uint) (*produk.Amount, error) {
	// panic("unimplemented")
	// mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	data, err := query.FindAmount(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetAmouts implements produk.Manager
func (m *manager) GetAmouts(ctx context.Context) ([]*produk.Amount, error) {
	query := m.repo.NewQuery()
	data, err := query.FindAmounts(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// GetProduct implements produk.Manager
func (m *manager) GetProduct(ctx context.Context, id uint) (*produk.Product, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProduct(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetProductType implements produk.Manager
func (m *manager) GetProductType(ctx context.Context, id uint) (*produk.ProductType, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProductType(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetProductTypes implements produk.Manager
func (m *manager) GetProductTypes(ctx context.Context) ([]*produk.ProductType, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProductTypes(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetProducts implements produk.Manager
func (m *manager) GetProducts(ctx context.Context) ([]*produk.Product, error) {
	query := m.repo.NewQuery()
	data, err := query.FindProducts(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetSatuan implements produk.Manager
func (m *manager) GetSatuan(ctx context.Context, id uint) (*produk.Satuan, error) {
	query := m.repo.NewQuery()
	data, err := query.FindSatuan(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// GetSatuans implements produk.Manager
func (m *manager) GetSatuans(ctx context.Context) ([]*produk.Satuan, error) {
	query := m.repo.NewQuery()
	data, err := query.FindSatuans(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateAmout implements produk.Manager
func (m *manager) UpdateAmout(ctx context.Context, a produk.Amount) (*produk.Amount, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.UpdateAmout(ctx, a)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindAmount(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateProduct implements produk.Manager
func (m *manager) UpdateProduct(ctx context.Context, p produk.UpdateProductRequest) (*produk.Product, error) {
	// panic("unimplemented")
	userId, _ := ctxutil.GetUserIdFromCtx(ctx)
	mutation := m.repo.NewMutation(ctx)
	product := produk.UpdateProductt(p, *userId)
	query := m.repo.NewQuery()
	id, err := mutation.UpdateProduct(ctx, product)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindProduct(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// UpdateProductType implements produk.Manager
func (m *manager) UpdateProductType(ctx context.Context, p produk.ProductType) (*produk.ProductType, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.UpdateProductType(ctx, p)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindProductType(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// UpdateSatuan implements produk.Manager
func (m *manager) UpdateSatuan(ctx context.Context, s produk.Satuan) (*produk.Satuan, error) {
	// panic("unimplemented")
	mutation := m.repo.NewMutation(ctx)
	query := m.repo.NewQuery()
	id, err := mutation.UpdateSatuan(ctx, s)
	if err != nil {
		log.Println(err)
		mutation.Cancel(ctx)
		return nil, err
	}
	if err = mutation.Commit(ctx); err != nil {
		return nil, err
	}
	data, err := query.FindSatuan(ctx, *id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
