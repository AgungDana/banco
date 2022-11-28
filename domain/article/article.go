package article

import "context"

type ArticleManager interface {
	GetArticle(ctx context.Context, id uint) (*ArticleResponse, error)
	// GetArticle(ctx context.Context) (*Artikel, error)
	GetListArticle(ctx context.Context) ([]*ArticleResponse, error)
}

type ArticleRepo interface {
	Create(Artikel) (*uint, error)
	FindById(uint) (*Artikel, error)
	FindAll() ([]*Artikel, error)
	Update(Artikel) (*uint, error)
	Delete(uint) error
}
