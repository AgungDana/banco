package route

import (
	"banco/httproute/service"

	"github.com/gin-gonic/gin"
)

type route struct {
	s service.Serv
	g *gin.Engine
}

func NewRoute(s service.Serv, g *gin.Engine) {
	a := route{
		s: s,
		g: g,
	}
	a.auth()
	a.articelRoute()
	a.noAuth()
}
func (r *route) articelRoute() {

}
func (r *route) auth() {
	//sedikit paham
	r.g.GET("get_user", func(ctx *gin.Context) {
		val := ctx.GetHeader("Authorization")
		if val == "" {
			ctx.JSON(400, gin.H{
				"messages": "need auth",
			})
			ctx.Abort()
			return
		}
	},
		func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"messages": "successs",
			})
		})
}

func (r *route) noAuth() {
	r.g.POST("get_article", r.s.GetArticle)
	r.g.GET("get_list_article", r.s.GetListArticle)
}
