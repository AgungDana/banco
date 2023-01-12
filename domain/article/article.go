package article

import "context"

type ArticleManager interface {
	CreateArticle(ctx context.Context, in CreateArticleRequest) (*ArticleResponse, error)
	DeleteArticle(ctx context.Context, in ArticleRequest) (*uint, error)
	UpdateArticle(ctx context.Context, in UpdateArticleRequest) (*ArticleResponse, error)
	GetArticle(ctx context.Context, in uint) (*ArticleResponse, error)
	GetListArticle(ctx context.Context) ([]*ArticleResponse, error)
}

type ArticleRepo interface {
	Create(Article) (*uint, error)
	FindById(uint) (*Article, error)
	FindAll() ([]*Article, error)
	Update(Article) (*uint, error)
	DeleteById(uint) (*uint, error)
}
