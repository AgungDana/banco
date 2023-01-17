package route

import (
	"banco/common/config"
	"banco/common/ctxutil"
	"banco/httproute/service"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type route struct {
	s    service.Service
	conf config.Config
}

func NewRoute(conf config.Config, s service.Service, g *gin.Engine) {
	a := route{
		s:    s,
		conf: conf,
	}
	g.Use(ctxutil.SetTransaction)
	a.noAuth(g)
	a.auth(g)
	a.ImageRouth(g)
}

func (r *route) ImageRouth(g *gin.Engine) {
	g.Use(static.Serve("/assets/image/product", static.LocalFile(r.conf.ImagePath, false)))
}

func (r *route) auth(g *gin.Engine) {
	gGroup := g.Group("", ctxutil.GinAuth)
	//user
	gGroup.GET("/profile", r.s.Profile)
	gGroup.PUT("/update-profile", r.s.UpdateUser)

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

	//produk
	gGroup.POST("/get-produk", r.s.GetProduct)
	gGroup.GET("/get-list-produk", r.s.GetProducts)
	gGroup.POST("/create-produk", r.s.CreateProduct)
	gGroup.PUT("/update-produk", r.s.UpdateProduct)
	gGroup.POST("/delete-produk", r.s.DeleteProduct)

	//tipe produk
	gGroup.POST("/get-produk-type", r.s.GetProductType)
	gGroup.GET("/get-list-produk-type", r.s.GetProductTypes)
	gGroup.POST("/create-produk-type", r.s.CreateProductType)
	gGroup.PUT("/update-produk-type", r.s.UpdateProductType)
	gGroup.POST("/delete-produk-type", r.s.DeleteProductType)

	//total
	gGroup.POST("/get-total", r.s.GetAmout)
	gGroup.GET("/get-list-total", r.s.GetAmouts)
	gGroup.POST("/create-total", r.s.CreateAmout)
	gGroup.POST("/update-total", r.s.UpdateAmout)
	gGroup.POST("/delete-total", r.s.DeleteAmout)

	//satuan
	gGroup.POST("/get-satuan", r.s.GetSatuan)
	gGroup.GET("/get-list-satuan", r.s.GetSatuans)
	gGroup.POST("/create-satuan", r.s.CreateSatuan)
	gGroup.POST("/update-satuan", r.s.UpdateSatuan)
	gGroup.POST("/delete-satuan", r.s.DeleteSatuan)

	//address
	gGroup.POST("/get-list-province", r.s.GetListProvince)
	gGroup.POST("/get-list-city", r.s.GetListCities)
	gGroup.POST("/get-list-district", r.s.GetListDistrict)

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

	g.POST("/upload-image-product", r.s.SaveProductImage)
}
