package manager

import (
	"banco/common/config"
	"banco/common/ctxutil"
	"banco/domain/article"
	"banco/domain/article/repo"
	"context"
	"fmt"
	"log"
)

func NewArticleManager(conf config.Config) article.ArticleManager {
	r, err := repo.NewRepoArticle(conf)
	if err != nil {
		// panic(err)
		log.Fatal(err)
		return nil
	}
	return &manager{
		repo: r,
	}
}

type manager struct {
	repo article.ArticleRepo
}

// CreateArticle implements article.ArticleManager
func (m *manager) CreateArticle(ctx context.Context, in article.CreateArticleRequest) (*article.ArticleResponse, error) {
	payload, _ := ctxutil.GetUserPayloadFromCtx(ctx)
	id, err := m.repo.Create(in.NewArticle(payload.UserID))
	if err != nil {
		return nil, err
	}

	data, err := m.repo.FindById(*id)
	if err != nil {
		return nil, err
	}
	return article.ToResponse(data), nil
}

// DeleteArticle implements article.ArticleManager
func (m *manager) DeleteArticle(ctx context.Context, in article.ArticleRequest) (*uint, error) {
	id, err := m.repo.DeleteById(in.ArticleId)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// UpdateArticle implements article.ArticleManager
func (m *manager) UpdateArticle(ctx context.Context, in article.UpdateArticleRequest) (*article.ArticleResponse, error) {
	payload, _ := ctxutil.GetUserPayloadFromCtx(ctx)
	id, err := m.repo.Update(in.ChangeArticle(payload.UserID))
	if err != nil {
		return nil, err
	}

	data, err := m.repo.FindById(*id)
	if err != nil {
		return nil, err
	}
	return article.ToResponse(data), nil
}

// GetArticle implements artikel.article
func (m *manager) GetArticle(ctx context.Context, in uint) (*article.ArticleResponse, error) {
	data, err := m.repo.FindById(in)
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

// GetListArticle implements artikel.ArticleManager
func (m *manager) GetListArticle(ctx context.Context) ([]*article.ArticleResponse, error) {
	data, err := m.repo.FindAll()
	if err != nil {
		fmt.Println("failed get list article")
		return nil, err
	}

	res := []*article.ArticleResponse{}
	for _, v := range data {
		res = append(res, &article.ArticleResponse{
			Id:          v.Id,
			Title:       v.Title,
			Description: v.Description,
			TimeStart:   v.TimeStart,
			TimeEnd:     v.TimeEnd,
		})
	}
	return res, nil
}
