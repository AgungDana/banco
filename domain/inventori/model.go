package inventori

import (
	"banco/common/infra/orm"
	"banco/domain/produk"
	"time"
)

type Inventory struct {
	orm.Model
	ProductId     uint           `json:"productId,omitempty"`
	Product       produk.Product `json:"product,omitempty"`
	IncomingItem  int            `json:"incomingItem,omitempty"`
	OutcomingItem int            `json:"outcomingItem,omitempty"`
	Stock         int            `json:"stock,omitempty"`
	Balance       int            `json:"balance,omitempty"`
	SatuanId      int            `json:"satuanId,omitempty"`
}

type CreateRequest struct {
	ProductId    uint `json:"productId,omitempty"`
	IncomingItem int  `json:"incomingItem,omitempty"`
	SatuanId     int  `json:"satuanId,omitempty"`
}

type InventoryIncreaseRequest struct {
	Id           uint `gorm:"primarykey" json:"id,omitempty"`
	ProductId    uint `json:"productId,omitempty"`
	IncomingItem int  `json:"incomingItem,omitempty"`
	Balance      int  `json:"balance,omitempty"`
	SatuanId     int  `json:"satuanId,omitempty"`
}

func IncreaseReqToInventory(req InventoryIncreaseRequest) Inventory {
	return Inventory{
		Model: orm.Model{
			Id: req.Id,
		},
		ProductId:    req.ProductId,
		IncomingItem: req.IncomingItem,
		Balance:      req.Balance,
		SatuanId:     req.SatuanId,
	}
}

type InventoryDecreaseRequest struct {
	Id            uint `gorm:"primarykey" json:"id,omitempty"`
	ProductId     uint `json:"productId,omitempty"`
	OutcomingItem int  `json:"outcomingItem,omitempty"`
	Balance       int  `json:"balance,omitempty"`
	SatuanId      int  `json:"satuanId,omitempty"`
}

func DecreaseReqToInventory(req InventoryDecreaseRequest) Inventory {
	return Inventory{
		Model: orm.Model{
			Id: req.Id,
		},
		ProductId:     req.ProductId,
		OutcomingItem: req.OutcomingItem,
		Balance:       req.Balance,
		SatuanId:      req.SatuanId,
	}
}

type IncomingItem struct {
	orm.Model
	TransactionId uint      `json:"transactionId,omitempty"`
	ProductId     int       `json:"productId,omitempty"`
	TotalItem     int       `json:"totalItem,omitempty"`
	Date          time.Time `json:"date,omitempty"`
}

type HistoryOutcomingItem struct {
	orm.Model
	TransactionId uint      `json:"transactionId,omitempty"`
	ProductId     int       `json:"productId,omitempty"`
	TotalItem     int       `json:"totalItem,omitempty"`
	Date          time.Time `json:"date,omitempty"`
}
