package article

import "context"

type ArticleManager interface {
	GetArticle(ctx context.Context, id uint) (*ArticleResponse, error)
	// GetArticle(ctx context.Context) (*Artikel, error)
	GetListArticle(ctx context.Context) ([]*ArticleResponse, error)
}

type ArticleRepo interface {
	Create(Article) (*uint, error)
	FindById(uint) (*Article, error)
	FindAll() ([]*Article, error)
	Update(Article) (*uint, error)
	Delete(uint) error
}
