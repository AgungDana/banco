package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/inventori"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepoInventory(conf config.Config) (inventori.Repo, error) {
	if conf.Driver != "" {
		log.Println("creating gorm repo")
		return newGormRepo(conf)
	}
	return nil, errors.New("driver not set")
}

func newGormRepo(conf config.Config) (inventori.Repo, error) {
	db, err := orm.DbCon(conf)
	if err != nil {
		return nil, err
	}
	return &repo{db: db}, nil
}

// NewMutation implements inventori.Repo
func (r *repo) NewMutation(ctx context.Context) inventori.Mutation {
	return newMutation(ctx, r.db)
}

// NewQuery implements inventori.Repo
func (r *repo) NewQuery() inventori.Query {
	return newQuery(r.db)
}
