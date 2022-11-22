package manager

import (
	"banco/common/config"
	"banco/domain/artikel"
	"banco/domain/artikel/repo"
	"context"
)

func NewManager(conf config.Config) artikel.ArticleManager {
	repo.NewRepoArticle(conf)
	return &manager{}
}

type manager struct {
}

// GetArticle implements artikel.ArticleManager
func (*manager) GetArticle(ctx context.Context, id string) (*artikel.Artikel, error) {
	panic("unimplemented")
}

// GetListArticle implements artikel.ArticleManager
func (*manager) GetListArticle(ctx context.Context, id string) (*artikel.Artikel, error) {
	panic("unimplemented")
}
