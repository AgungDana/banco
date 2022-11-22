package route

import (
	"banco/httproute/route/service"

	"github.com/gin-gonic/gin"
)

type ArticleRoute struct {
	s service.Service
	r *gin.Engine
}

func NewAritcleRoute(s service.Service, r *gin.Engine) *gin.Engine {
	a := ArticleRoute{
		s: s,
		r: r,
	}

	a.r.GET("get_article", a.GetArticle)

	return a.r
}

func (a *ArticleRoute) GetArticle(c *gin.Context) {

}
