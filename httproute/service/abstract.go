package service

import "github.com/gin-gonic/gin"

type Service interface {
	GetArticle(c *gin.Context)
	GetListArticle(c *gin.Context)
	CreateArticle(c *gin.Context)
	UpdateArticle(c *gin.Context)
	DeleteArticle(c *gin.Context)

	Login(c *gin.Context)
	Profile(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	FindAllInventory(c *gin.Context)
	FindAInventoryByIdInventory(c *gin.Context)
	FindAInventoryByIdProduct(c *gin.Context)
	AddProductToInvetory(c *gin.Context)
	IncreaseProductToInventory(c *gin.Context)
	DecreaseProductToInventory(c *gin.Context)
	DeleteProductFromInventory(c *gin.Context)

	FindTransaksiByID(c *gin.Context)
	FindTransaksiByProductID(c *gin.Context)
	FindAllTransaksi(c *gin.Context)
	CreateTransaksi(c *gin.Context)

	//Produk
	GetProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)

	GetProductType(c *gin.Context)
	GetProductTypes(c *gin.Context)
	CreateProductType(c *gin.Context)
	UpdateProductType(c *gin.Context)
	DeleteProductType(c *gin.Context)

	GetAmout(c *gin.Context)
	GetAmouts(c *gin.Context)
	CreateAmout(c *gin.Context)
	UpdateAmout(c *gin.Context)
	DeleteAmout(c *gin.Context)

	GetSatuan(c *gin.Context)
	GetSatuans(c *gin.Context)
	CreateSatuan(c *gin.Context)
	UpdateSatuan(c *gin.Context)
	DeleteSatuan(c *gin.Context)

	GetListProvince(c *gin.Context)
	GetListCities(c *gin.Context)
	GetListDistrict(c *gin.Context)

	SaveProductImage(c *gin.Context)
}
