package produk

import "banco/common/infra/orm"

type ProductType struct {
	orm.Model
	Type          string
	Desccriptions string
}
