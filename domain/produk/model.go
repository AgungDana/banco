package produk

import "banco/common/infra/orm"

type ProductType struct {
	orm.Model
	Type          string
	Desccriptions string
}

type Product struct {
	orm.Model
	TypeProductId int
	SpecificType  string
	Description   string
}

type ProductDiscount struct {
	orm.Model
	DiscountId int
	ProductId  int
}
