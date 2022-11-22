package article

import "context"

type ArticleManager interface {
	GetArticle(ctx context.Context, id uint) (*ArticleResponse, error)
	GetListArticle(ctx context.Context) ([]*ArticleResponse, error)
}

type ArticleRepo interface {
	Create(Artikel) (*uint, error)
	FindById(uint) (*Artikel, error)
	FindAll() ([]*Artikel, error)
	Update(Artikel) (*uint, error)
	Delete(uint) error
}
