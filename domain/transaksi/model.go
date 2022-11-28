package transaksi

import (
	"banco/common/infra/orm"
	"time"
)

type Transaction struct {
	orm.Model
	TotalTransaction int
	UserId           int
	UserName         string
	DateTransactions time.Time
	Amount           int
}

type TransactionDetail struct {
	orm.Model
	TransactionId int
	ProductTypeId int
	ProductName   string
	AmountId      int
	DiscountId    int
	Discount      int
	Amount        int
	TotalItem     int
	Cost          int
}

type Amount struct {
	orm.Model
	ProductId int
	SatuanId  int
	Cost      int
}

type Satuan struct {
	orm.Model
	Name string
}
