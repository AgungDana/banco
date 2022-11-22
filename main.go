package main

import (
	"banco/common/config"
	"banco/httproute/route"
	"banco/httproute/route/service"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()

	r := gin.Default()
	as := service.NewServiceArticle(conf)
	route.NewAritcleRoute(as, r)

	r.Run()

}
