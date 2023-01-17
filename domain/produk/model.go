package produk

import "banco/common/infra/orm"

type Product struct {
	orm.Model
	ProductTypeId uint        `json:"productTypeId,omitempty"`
	Image         string      `json:"image,omitempty"`
	ProductType   ProductType `json:"productType,omitempty"`
	SpecificType  string      `json:"specificType,omitempty"`
	Description   string      `json:"description,omitempty"`
}

type CreateProductRequest struct {
	ProductTypeId uint        `json:"productTypeId,omitempty"`
	ProductType   ProductType `json:"productType,omitempty"`
	SpecificType  string      `json:"specificType,omitempty"`
	Description   string      `json:"description,omitempty"`
}

func NewProduct(req CreateProductRequest, userId uint) Product {
	return Product{
		Model: orm.Model{
			CraetedBy: userId,
		},
		ProductTypeId: req.ProductTypeId,
		ProductType:   req.ProductType,
		SpecificType:  req.SpecificType,
		Description:   req.Description,
	}
}

type UpdateProductRequest struct {
	ProductTypeId uint        `json:"productTypeId,omitempty"`
	ProductType   ProductType `json:"productType,omitempty"`
	SpecificType  string      `json:"specificType,omitempty"`
	Description   string      `json:"description,omitempty"`
}

func UpdateProduct(req UpdateProductRequest, userId uint) Product {
	return Product{
		Model: orm.Model{
			Id:        req.ProductTypeId,
			UpdatedBy: userId,
		},
		ProductTypeId: req.ProductTypeId,
		ProductType:   req.ProductType,
		SpecificType:  req.SpecificType,
		Description:   req.Description,
	}
}

type ProductType struct {
	orm.Model
	Type          string    `json:"type,omitempty"`
	Product       []Product `json:"product,omitempty"`
	Desccriptions string    `json:"desccriptions,omitempty"`
}

type CreateProductTypeRequest struct {
	Type          string    `json:"type,omitempty"`
	Product       []Product `json:"product,omitempty"`
	Desccriptions string    `json:"desccriptions,omitempty"`
}

func NewProductType(req CreateProductTypeRequest, userId uint) ProductType {
	return ProductType{
		Model: orm.Model{
			CraetedBy: userId,
		},
		Type:          req.Type,
		Product:       req.Product,
		Desccriptions: req.Desccriptions,
	}
}

type UpdateProductTypeRequest struct {
	Id            uint      `json:"id,omitempty"`
	Type          string    `json:"type,omitempty"`
	Product       []Product `json:"product,omitempty"`
	Desccriptions string    `json:"desccriptions,omitempty"`
}

func UpdateProductType(req UpdateProductTypeRequest, id uint) ProductType {
	return ProductType{
		Model: orm.Model{
			Id:        req.Id,
			UpdatedBy: id,
		},
		Type:          req.Type,
		Product:       req.Product,
		Desccriptions: req.Desccriptions,
	}
}

type ProductDiscount struct {
	orm.Model
	DiscountId uint `json:"discountId,omitempty"`
	ProductId  uint `json:"productId,omitempty"`
}

type CreateProductDiscountRequest struct {
	DiscountId uint `json:"discountId,omitempty"`
	ProductId  uint `json:"productId,omitempty"`
}

func NewProductDiscount(req CreateProductDiscountRequest, userId uint) ProductDiscount {
	return ProductDiscount{
		Model: orm.Model{
			CraetedBy: userId,
		},
		DiscountId: req.DiscountId,
		ProductId:  req.ProductId,
	}
}

type UpdateProductDiscountRequest struct {
	DiscountId uint
	ProductId  uint
}

// wertewrtw

type Amount struct {
	orm.Model
	ProductId uint
	SatuanId  uint
	Cost      int
}

type Satuan struct {
	orm.Model
	Name string
}
