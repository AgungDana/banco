package transaksi

import (
	"banco/common/infra/orm"
	"time"
)

type Transaction struct {
	orm.Model
	TotalTransaction  int
	TransactionDetail []TransactionDetail
	UserId            uint
	UserName          string
	DateTransactions  time.Time
	Amount            int
}

type TransactionReq struct {
	Id                uint
	TotalTransaction  int
	TransactionDetail []TransactionDetailReq
	UserName          string
	DateTransactions  time.Time
	Amount            int
}

func ToTransaction(req TransactionReq, userId uint) Transaction {
	return Transaction{
		Model: orm.Model{
			CraetedBy: userId,
		},
		TotalTransaction:  req.TotalTransaction,
		TransactionDetail: ToTransactionDetail(req.TransactionDetail, userId),
		UserId:            userId,
		UserName:          req.UserName,
		DateTransactions:  req.DateTransactions,
		Amount:            req.Amount,
	}
}

func ToTransactionDetail(req []TransactionDetailReq, userid uint) []TransactionDetail {
	detail := []TransactionDetail{}
	for _, v := range req {
		detail = append(detail, TransactionDetail{
			Model: orm.Model{
				CraetedBy: userid,
			},
			ProductId:   v.ProductId,
			ProductName: v.ProductName,
			AmountId:    v.AmountId,
			DiscountId:  v.DiscountId,
			Discount:    v.Discount,
			Amount:      v.Amount,
			TotalItem:   v.TotalItem,
			Cost:        v.Cost,
		})
	}
	return detail
}

type TransactionDetail struct {
	orm.Model
	TransactionId uint
	Transaction   Transaction
	ProductId     uint
	ProductName   string
	AmountId      int
	DiscountId    int
	Discount      int
	Amount        int
	TotalItem     int
	Cost          int
}

type TransactionDetailReq struct {
	ProductId   uint
	ProductName string
	AmountId    int
	DiscountId  int
	Discount    int
	Amount      int
	TotalItem   int
	Cost        int
}
