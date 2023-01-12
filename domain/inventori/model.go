package inventori

import (
	"banco/common/infra/orm"
	"banco/domain/produk"
	"time"
)

type Inventory struct {
	orm.Model
	ProductId     uint
	Product       produk.Product
	IncomingItem  int
	OutcomingItem int
	Stock         int
	Balance       int
	SatuanId      int
}

type CreateRequest struct {
	ProductId    uint
	IncomingItem int
	Stock        int
	SatuanId     int
}

type InventoryIncreaseRequest struct {
	Id           uint `gorm:"primarykey" json:"id,omitempty"`
	ProductId    uint
	IncomingItem int
	Stock        int
	Balance      int
	SatuanId     int
}

func IncreaseReqToInventory(req InventoryIncreaseRequest) Inventory {
	return Inventory{
		Model: orm.Model{
			Id: req.Id,
		},
		ProductId:    req.ProductId,
		IncomingItem: req.IncomingItem,
		Stock:        req.Stock,
		Balance:      req.Balance,
		SatuanId:     req.SatuanId,
	}
}

type InventoryDecreaseRequest struct {
	Id            uint `gorm:"primarykey" json:"id,omitempty"`
	ProductId     uint
	OutcomingItem int
	Stock         int
	Balance       int
	SatuanId      int
}

func DecreaseReqToInventory(req InventoryDecreaseRequest) Inventory {
	return Inventory{
		Model: orm.Model{
			Id: req.Id,
		},
		ProductId:     req.ProductId,
		OutcomingItem: req.OutcomingItem,
		Stock:         req.Stock,
		Balance:       req.Balance,
		SatuanId:      req.SatuanId,
	}
}

type IncomingItem struct {
	orm.Model
	TransactionId uint
	ProductId     int
	TotalItem     int
	Date          time.Time
}

type HistoryOutcomingItem struct {
	orm.Model
	TransactionId uint
	ProductId     int
	TotalItem     int
	Date          time.Time
}
