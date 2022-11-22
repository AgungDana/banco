package service

import (
	"banco/common/config"
	"banco/domain/artikel"
	"banco/domain/artikel/manager"
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
	m artikel.ArticleManager
}
