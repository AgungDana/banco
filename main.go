package main

import (
	"banco/common/config"
	"banco/httproute/route"
	"banco/httproute/service"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()

	r := gin.Default()
	s := service.NewService(conf)
	route.NewRoute(s, r)

	r.Run()

}
