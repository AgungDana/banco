package service

import "github.com/gin-gonic/gin"

type Serv interface {
	GetArticle(c *gin.Context)
	GetListArticle(c *gin.Context)
}

// type Serv2 interface {
// 	GetArticle(c *gin.Context)
// 	GetListArticle(c *gin.Context)
// }
