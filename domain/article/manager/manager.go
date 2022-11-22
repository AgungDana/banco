package manager

import (
	"banco/common/config"
	"banco/domain/article"
	"banco/domain/article/repo"
	"context"
)

func NewManager(conf config.Config) article.ArticleManager {
	repo.NewRepoArticle(conf)
	return &manager{}
}

type manager struct {
}

// GetArticle implements artikel.article
func (*manager) GetArticle(ctx context.Context, id uint) (*article.ArticleResponse, error) {
	panic("unimplemented")
}

// GetListArticle implements artikel.ArticleManager
func (*manager) GetListArticle(ctx context.Context) ([]*article.ArticleResponse, error) {
	panic("unimplemented")
}
