package service

import (
	"banco/common/config"
	"banco/common/restsvr"
	"banco/domain/article"
	mArticle "banco/domain/article/manager"
	"banco/domain/inventori"
	mInventory "banco/domain/inventori/manager"
	"banco/domain/produk"
	mProduk "banco/domain/produk/manager"
	"banco/domain/transaksi"
	mTransaksi "banco/domain/transaksi/manager"
	"banco/domain/user"
	mUser "banco/domain/user/manager"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewService(conf config.Config) Service {
	m := mArticle.NewArticleManager(conf)
	u := mUser.NewUserManager(conf)
	i := mInventory.NewManager(conf)
	t := mTransaksi.NewManager(conf)
	p := mProduk.NewManager(conf)

	return &service{
		mArticle:   m,
		mUser:      u,
		mInventory: i,
		mTransaksi: t,
		mProduk:    p,
		conf:       conf,
	}
}

type service struct {
	mArticle   article.ArticleManager
	mUser      user.Manager
	mInventory inventori.Manager
	mTransaksi transaksi.Manager
	mProduk    produk.Manager
	conf       config.Config
}

// SaveProductImage implements Service
func (s *service) SaveProductImage(c *gin.Context) {

	type Images struct {
		Url   string `json:"url,omitempty"`
		Image string `json:"image,omitempty"`
	}
	res := new(restsvr.Repsonse)
	defer restsvr.CreateResponse(c, res)
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		c.Abort()
		return
	}
	ext := filepath.Ext(file.Filename)
	filename := uuid.New().String() + ext
	fmt.Println(filename)
	err = c.SaveUploadedFile(file, s.conf.ImagePath+"/"+filename)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data := Images{
		Url:   s.conf.ImageUrl + "/" + filename,
		Image: filename,
	}
	res.Add(&data, err)
}

// GetListCities implements Service
func (s *service) GetListCities(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = user.City{}
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mUser.GetListCity(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetListDistrict implements Service
func (s *service) GetListDistrict(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = user.District{}
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mUser.GetListDistrict(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetListProvince implements Service
func (s *service) GetListProvince(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = user.Country{}
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mUser.GetListProvince(c.Request.Context(), req.Id)

	res.Add(data, err)
}

// CreateAmout implements Service
func (s *service) CreateAmout(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = produk.Amount{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.CreateAmout(c.Request.Context(), req)
	res.Add(data, err)
}

// CreateProduct implements Service
func (s *service) CreateProduct(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.CreateProductRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	c.FormFile("image")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.CreateProduct(c.Request.Context(), req)
	res.Add(data, err)
}

// CreateProductType implements Service
func (s *service) CreateProductType(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = produk.ProductType{}
	)
	defer restsvr.CreateResponse(c, res)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.CreateProductType(c.Request.Context(), req)
	res.Add(data, err)
}

// CreateSatuan implements Service
func (s *service) CreateSatuan(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Satuan{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.CreateSatuan(c.Request.Context(), req)
	res.Add(data, err)
}

// DeleteAmout implements Service
func (s *service) DeleteAmout(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Amount{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.DeleteAmout(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// DeleteProduct implements Service
func (s *service) DeleteProduct(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Product{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.DeleteProduct(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// DeleteProductType implements Service
func (s *service) DeleteProductType(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.ProductType{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.DeleteProductType(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// DeleteSatuan implements Service
func (s *service) DeleteSatuan(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Satuan{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.DeleteSatuan(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetAmout implements Service
func (s *service) GetAmout(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Amount{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.GetAmout(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetAmouts implements Service
func (s *service) GetAmouts(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	res.Add(nil, err)
	// 	return
	// }
	data, err := s.mProduk.GetAmouts(c.Request.Context())
	res.Add(data, err)
}

// GetProduct implements Service
func (s *service) GetProduct(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Product{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.GetProduct(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetProductType implements Service
func (s *service) GetProductType(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.ProductType{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.GetProductType(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetProductTypes implements Service
func (s *service) GetProductTypes(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	res.Add(nil, err)
	// 	return
	// }
	data, err := s.mProduk.GetProductTypes(c.Request.Context())
	res.Add(data, err)
}

// GetProducts implements Service
func (s *service) GetProducts(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	res.Add(nil, err)
	// 	return
	// }
	data, err := s.mProduk.GetProducts(c.Request.Context())
	res.Add(data, err)
}

// GetSatuan implements Service
func (s *service) GetSatuan(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Product{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.GetSatuan(c.Request.Context(), req.Id)
	res.Add(data, err)
}

// GetSatuans implements Service
func (s *service) GetSatuans(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	// err := c.ShouldBindJSON(&req)
	// if err != nil {
	// 	res.Add(nil, err)
	// 	return
	// }
	data, err := s.mProduk.GetSatuans(c.Request.Context())
	res.Add(data, err)
}

// UpdateAmout implements Service
func (s *service) UpdateAmout(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Amount{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.UpdateAmout(c.Request.Context(), req)
	res.Add(data, err)
}

// UpdateProduct implements Service
func (s *service) UpdateProduct(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.UpdateProductRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.UpdateProduct(c.Request.Context(), req)
	res.Add(data, err)
}

// UpdateProductType implements Service
func (s *service) UpdateProductType(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.ProductType{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.UpdateProductType(c.Request.Context(), req)
	res.Add(data, err)
}

// UpdateSatuan implements Service
func (s *service) UpdateSatuan(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = produk.Satuan{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mProduk.UpdateSatuan(c.Request.Context(), req)
	res.Add(data, err)
}

// AddProductToInvetory implements Service
func (s *service) AddProductToInvetory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.CreateRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.AddProductToInventory(c.Request.Context(), req)
	res.Add(data, err)
}

// DecreaseProductToInventory implements Service
func (s *service) DecreaseProductToInventory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.InventoryDecreaseRequest{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.DecreaseProductInInventory(ctx, req)
	res.Add(data, err)
}

// DeleteProductFromInventory implements Service
func (s *service) DeleteProductFromInventory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.Inventory{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.DeleteProductFromInvetory(ctx, req.Id)
	res.Add(data, err)
}

// FindAInventoryByIdInventory implements Service
func (s *service) FindAInventoryByIdInventory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.Inventory{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.FindProductFromInventoryByIdInventory(ctx, req.Id)
	res.Add(data, err)
}

// FindAInventoryByIdProduct implements Service
func (s *service) FindAInventoryByIdProduct(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.Inventory{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.FindProductFromInventoryByIdProduct(ctx, req.ProductId)
	res.Add(data, err)
}

// FindAllInventory implements Service
func (s *service) FindAllInventory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.Inventory{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mInventory.FindAll(ctx)
	res.Add(data, err)
}

// IncreaseProductToInventory implements Service
func (s *service) IncreaseProductToInventory(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = inventori.InventoryIncreaseRequest{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mInventory.IncreaseProductInInventory(ctx, req)
	res.Add(data, err)
}

// // UpdateProductFromInventory implements Service
// func (s *service) UpdateProductFromInventory(c *gin.Context) {
// 	var (
// 		res = new(restsvr.Repsonse)
// 		req = inventori.Inventory{}
// 		ctx = c.Request.Context()
// 	)
// 	defer restsvr.CreateResponse(c, res)
// 	c.ShouldBindJSON(&req)
// 	data, err := s.mInventory.
// 		res.Add(data, err)
// 	panic("unimplemented")
// }

// Login implements Service
func (s *service) Login(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = user.LoginRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	// err := c.ShouldBind(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mUser.Login(c.Request.Context(), req)
	res.Add(data, err)
}

// Profile implements Service
func (s *service) Profile(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	data, err := s.mUser.FindUserByID(c.Request.Context())
	res.Add(data, err)
}

// CreateUser implements Service
func (s *service) CreateUser(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = user.CreateUserRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	// data := map[string]any{}

	// json.NewDecoder(c.Request.Body).Decode(&data)
	// fmt.Println(data)
	err := c.ShouldBindJSON(&req)
	// err := c.ShouldBind(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mUser.CreateUser(c.Request.Context(), req)
	res.Add(data, err)
}

// DeleteUser implements Service
func (s *service) DeleteUser(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
	)
	defer restsvr.CreateResponse(c, res)
	data, err := s.mUser.DeleteUser(c.Request.Context())
	res.Add(data, err)
}

// UpdateUser implements Service
func (s *service) UpdateUser(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = user.UpdateUserRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	fmt.Printf("req: %v\n", req)
	err := c.ShouldBindJSON(&req)
	fmt.Printf("req2: %v\n", req)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		res.Add(nil, err)
		return
	}
	data, err := s.mUser.UpdateUserById(c.Request.Context(), req)
	fmt.Printf("erre: %v\n", err)
	res.Add(data, err)
}

// CreateArticle implements Service
func (s *service) CreateArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.CreateArticleRequest
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Print("Errror ")
		fmt.Print(err)
		return
	}
	defer restsvr.CreateResponse(c, res)
	data, err := s.mArticle.CreateArticle(c.Request.Context(), req)
	res.Add(data, err)
}

// DeleteArticle implements Service
func (s *service) DeleteArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.ArticleRequest
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	defer restsvr.CreateResponse(c, res)
	data, err := s.mArticle.DeleteArticle(c.Request.Context(), req)
	res.Add(data, err)
}

// UpdateArticle implements Service
func (s *service) UpdateArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.UpdateArticleRequest
	)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	defer restsvr.CreateResponse(c, res)
	data, err := s.mArticle.UpdateArticle(c.Request.Context(), req)
	res.Add(data, err)
}

func (s *service) GetArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.ArticleRequest
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mArticle.GetArticle(c.Request.Context(), req.ArticleId)
	res.Add(data, err)
}

func (s *service) GetListArticle(c *gin.Context) {
	res := new(restsvr.Repsonse)

	defer restsvr.CreateResponse(c, res)
	data, err := s.mArticle.GetListArticle(c.Request.Context())
	res.Add(data, err)
}

func (s *service) FindTransaksiByID(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = transaksi.TransactionReq{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mTransaksi.FindTrasaksiByID(ctx, req.Id)
	res.Add(data, err)
}
func (s *service) FindTransaksiByProductID(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = transaksi.TransactionReq{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mTransaksi.FindTrasaksiByProductID(ctx, req.Id)
	res.Add(data, err)
}
func (s *service) FindAllTransaksi(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	data, err := s.mTransaksi.FindAllTrasaksi(ctx)
	res.Add(data, err)
}
func (s *service) CreateTransaksi(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req = transaksi.TransactionReq{}
		ctx = c.Request.Context()
	)
	defer restsvr.CreateResponse(c, res)
	c.ShouldBindJSON(&req)
	data, err := s.mTransaksi.CreateTransaksi(ctx, req)
	res.Add(data, err)
}
