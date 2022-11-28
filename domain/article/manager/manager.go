package manager

import (
	"banco/common/config"
	"banco/domain/article"
	"banco/domain/article/repo"
	"context"
)

func NewManager(conf config.Config) article.ArticleManager {
	r := repo.NewRepoArticle(conf)
	return &manager{r: r}
}

type manager struct {
	r article.ArticleRepo
}

// GetArticle implements artikel.article
func (m *manager) GetArticle(ctx context.Context, id uint) (*article.ArticleResponse, error) {
	data, err := m.r.FindById(id)
	if err != nil {
		return nil, err
	}
	return &article.ArticleResponse{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		TimeStart:   data.TimeStart,
		TimeEnd:     data.TimeEnd,
	}, nil
}

// GetArticle implements artikel.article
// func (*manager) GetArticle(ctx context.Context) (*article.Artikel, error) {
// 	panic("unimplemented")
// }
// GetListArticle implements artikel.ArticleManager
func (*manager) GetListArticle(ctx context.Context) ([]*article.ArticleResponse, error) {
	panic("unimplemented")
}
