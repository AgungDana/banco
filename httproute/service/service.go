package service

import (
	"banco/common/config"
	"banco/common/restsvr"
	"banco/domain/article"
	"banco/domain/article/manager"

	"github.com/gin-gonic/gin"
)

func NewService(conf config.Config) Serv {
	m := manager.NewManager(conf)

	return &service{
		m: m,
	}
}

type service struct {
	m article.ArticleManager
}

func (s *service) GetArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.ArticleRequest
	)
	c.ShouldBindJSON(&req)
	defer restsvr.CreateResponse(c, res)
	// data, err := s.m.GetArticle(c)
	data, err := s.m.GetArticle(c, req.ArticleId)
	res.Add(data, err)
}

func (a *service) GetListArticle(c *gin.Context) {
}
