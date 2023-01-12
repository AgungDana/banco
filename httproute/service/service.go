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

	"github.com/gin-gonic/gin"
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
	}
}

type service struct {
	mArticle   article.ArticleManager
	mUser      user.Manager
	mInventory inventori.Manager
	mTransaksi transaksi.Manager
	mProduk    produk.Manager
}

// CreateAmout implements Service
func (s *service) CreateAmout(c *gin.Context) {
	// panic("unimplemented")
	var (
		res = new(restsvr.Repsonse)
		req = produk.CreateProductRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
}

// CreateProduct implements Service
func (s *service) CreateProduct(c *gin.Context) {
	panic("unimplemented")
}

// CreateProductType implements Service
func (s *service) CreateProductType(c *gin.Context) {
	panic("unimplemented")
}

// CreateSatuan implements Service
func (s *service) CreateSatuan(c *gin.Context) {
	panic("unimplemented")
}

// DeleteAmout implements Service
func (s *service) DeleteAmout(c *gin.Context) {
	panic("unimplemented")
}

// DeleteProduct implements Service
func (s *service) DeleteProduct(c *gin.Context) {
	panic("unimplemented")
}

// DeleteProductType implements Service
func (s *service) DeleteProductType(c *gin.Context) {
	panic("unimplemented")
}

// DeleteSatuan implements Service
func (s *service) DeleteSatuan(c *gin.Context) {
	panic("unimplemented")
}

// GetAmout implements Service
func (s *service) GetAmout(c *gin.Context) {
	panic("unimplemented")
}

// GetAmouts implements Service
func (s *service) GetAmouts(c *gin.Context) {
	panic("unimplemented")
}

// GetProduct implements Service
func (s *service) GetProduct(c *gin.Context) {
	panic("unimplemented")
}

// GetProductType implements Service
func (s *service) GetProductType(c *gin.Context) {
	panic("unimplemented")
}

// GetProductTypes implements Service
func (s *service) GetProductTypes(c *gin.Context) {
	panic("unimplemented")
}

// GetProducts implements Service
func (s *service) GetProducts(c *gin.Context) {
	panic("unimplemented")
}

// GetSatuan implements Service
func (s *service) GetSatuan(c *gin.Context) {
	panic("unimplemented")
}

// GetSatuans implements Service
func (s *service) GetSatuans(c *gin.Context) {
	panic("unimplemented")
}

// UpdateAmout implements Service
func (s *service) UpdateAmout(c *gin.Context) {
	panic("unimplemented")
}

// UpdateProduct implements Service
func (s *service) UpdateProduct(c *gin.Context) {
	panic("unimplemented")
}

// UpdateProductType implements Service
func (s *service) UpdateProductType(c *gin.Context) {
	panic("unimplemented")
}

// UpdateSatuan implements Service
func (s *service) UpdateSatuan(c *gin.Context) {
	panic("unimplemented")
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
	c.ShouldBindJSON(&req)
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
		req = user.CreateUserRequest{}
	)
	defer restsvr.CreateResponse(c, res)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
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

	err := c.ShouldBindJSON(&req)
	if err != nil {
		res.Add(nil, err)
		return
	}
	data, err := s.mUser.UpdateUserById(c.Request.Context(), req)
	res.Add(data, err)
}

// CreateArticle implements Service
func (s *service) CreateArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.CreateArticleRequest
	)
	c.ShouldBindJSON(&req)
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
	c.ShouldBindJSON(&req)
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
	c.ShouldBindJSON(&req)
	defer restsvr.CreateResponse(c, res)
	data, err := s.mArticle.UpdateArticle(c.Request.Context(), req)
	res.Add(data, err)
}

func (s *service) GetArticle(c *gin.Context) {
	var (
		res = new(restsvr.Repsonse)
		req article.ArticleRequest
	)
	c.ShouldBindJSON(&req)
	defer restsvr.CreateResponse(c, res)
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
