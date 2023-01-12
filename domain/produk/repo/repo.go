package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/produk"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

// NewMutation implements produk.Repo

func NewRepoProduct(conf config.Config) (produk.Repo, error) {
	if conf.Driver != "" {
		log.Println("cretating gorm repo")
		return newGormRepo(conf)
	}
	return nil, errors.New("driver not set")
}

func newGormRepo(conf config.Config) (produk.Repo, error) {
	db, err := orm.DbCon(conf)
	if err != nil {
		return nil, err
	}
	return &gormRepo{
		db: db,
	}, nil
}

func (g *gormRepo) NewMutation(ctx context.Context) produk.Mutation {
	return newMutation(ctx, g.db)
}

// NewQuery implements produk.Repo
func (g *gormRepo) NewQuery() produk.Query {
	return newQuery(g.db)
}
