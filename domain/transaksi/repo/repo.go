package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/inventori"
	"banco/domain/transaksi"
	"context"
	"errors"

	"gorm.io/gorm"
)

type repo struct {
	db           *gorm.DB
	invetoriRepo inventori.Repo
}

func NewRepo(conf config.Config, repoInventory inventori.Repo) (transaksi.Repo, error) {
	if conf.Driver != "" {
		return newRepo(conf, repoInventory)
	}
	return nil, errors.New("driver not set")
}

func newRepo(conf config.Config, repoInventory inventori.Repo) (transaksi.Repo, error) {
	db, err := orm.DbCon(conf)
	if err != nil {
		return nil, err
	}
	return &repo{db: db, invetoriRepo: repoInventory}, nil
}

func (r *repo) NewMutation(ctx context.Context) transaksi.Mutation {
	return newMutation(ctx, r.db, r.invetoriRepo)
}
func (r *repo) NewQuery() transaksi.Query {
	return newQuery(r.db)
}
