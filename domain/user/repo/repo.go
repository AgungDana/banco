package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/user"
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

type gormRepo struct {
	db *gorm.DB
}

func NewRepoUser(conf config.Config) (user.Repo, error) {
	if conf.Driver != "" {
		log.Println("creating gorm repo")
		return newGormRepo(conf)
	}
	return nil, errors.New("driver not set")
}

func newGormRepo(conf config.Config) (user.Repo, error) {
	db, err := orm.DbCon(conf)
	if err != nil {
		return nil, err
	}
	return &gormRepo{
		db: db,
	}, nil
}

// NewMutation implements user.Repo
func (g *gormRepo) NewMutation(ctx context.Context) user.Mutation {
	return newMutation(ctx, g.db)
}

// NewQuery implements user.Repo
func (g *gormRepo) NewQuery() user.Query {
	return newQuery(g.db)
}
