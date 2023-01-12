package route

import (
	"banco/common/ctxutil"
	"banco/httproute/service"

	"github.com/gin-gonic/gin"
)

type route struct {
	s service.Service
}

func NewRoute(s service.Service, g *gin.Engine) {
	a := route{
		s: s,
	}
	g.Use(ctxutil.SetTransaction)
	a.noAuth(g)
	a.auth(g)
}

func (r *route) auth(g *gin.Engine) {
	gGroup := g.Group("", ctxutil.GinAuth)
	//user
	gGroup.GET("/profile", r.s.Profile)
	gGroup.POST("/update-profile", r.s.UpdateUser)

	//article
	gGroup.POST("/create-article", r.s.CreateArticle)
	gGroup.POST("/update-article", r.s.UpdateArticle)
	gGroup.POST("/delete-article", r.s.DeleteArticle)

	//inventory
	gGroup.POST("/decrease-inventory", r.s.DecreaseProductToInventory)
	gGroup.POST("/increase-inventory", r.s.IncreaseProductToInventory)
	gGroup.POST("/delete-inventory", r.s.DeleteProductFromInventory)
	gGroup.POST("/add-inventory", r.s.AddProductToInvetory)
	gGroup.POST("/transaksi", r.s.CreateTransaksi)
}

func (r *route) noAuth(g *gin.Engine) {
	g.POST("/register", r.s.CreateUser)
	g.POST("/login", r.s.Login)

	g.POST("/get-article", r.s.GetArticle)
	g.GET("/get-list-article", r.s.GetListArticle)

	g.GET("/get-list-inventory", r.s.FindAllInventory)
	g.POST("/get-inventory-by-id-inventory", r.s.FindAInventoryByIdInventory)
	g.POST("/get-inventory-by-id-product", r.s.FindAInventoryByIdProduct)

	g.GET("/get-list-transaksi", r.s.FindAllTransaksi)
	g.POST("/get-transaksi-by-id", r.s.FindTransaksiByID)
	g.POST("/get-transaksi-by-product-id", r.s.FindTransaksiByProductID)
}
