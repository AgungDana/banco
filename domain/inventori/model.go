package inventori

import (
	"banco/common/infra/orm"
	"banco/domain/produk"
	"time"
)

type Inventory struct {
	orm.Model
	ProductTypeId int
	ProductType   produk.ProductType
	IncomingItem  int
	OutcomingItem int
	Stock         int
	Balance       int
	SatuanId      int
}

type HistoryOutcomingItem struct {
	orm.Model
	TransactionId int
	ProductId     int
	TotalItem     int
	Date          time.Time
}
