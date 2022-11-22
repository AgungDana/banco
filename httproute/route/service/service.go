package service

import (
	"banco/common/config"
	"banco/domain/article"
	"banco/domain/article/manager"
)

type Service interface {
}

func NewServiceArticle(conf config.Config) Service {
	m := manager.NewManager(conf)

	return &service{
		m: m,
	}
}

type service struct {
	m article.ArticleManager
}
