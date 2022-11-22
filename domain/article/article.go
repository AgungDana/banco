package article

import "context"

type ArticleManager interface {
	GetArticle(ctx context.Context, id string) (*Artikel, error)
	GetListArticle(ctx context.Context) (*Artikel, error)
}

type ArticleRepo interface {
	Create(Artikel) (*string, error)
	Read(string) (*string, error)
	Update(Artikel) (*string, error)
	Delete(string) (*string, error)
}
