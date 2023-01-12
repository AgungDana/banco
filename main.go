package main

import (
	"banco/common/config"
	"banco/httproute/route"
	"banco/httproute/service"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	r := gin.Default()
	r.Use(cors.Default())
	s := service.NewService(conf)
	route.NewRoute(s, r)
	//
	r.Run(conf.Address)
	// createTable()
}

// func createTable() {
// 	db, err := orm.DbCon(config.GetConfig())
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	db.AutoMigrate(
// 		article.Article{},
// 		diskon.Discount{},
// 		produk.ProductType{},
// 		produk.Product{},
// 		produk.ProductDiscount{},
// 		produk.Amount{},
// 		produk.Satuan{},
// 		inventori.Inventory{},
// 		inventori.HistoryOutcomingItem{},
// 		transaksi.Transaction{},
// 		transaksi.TransactionDetail{},
// 		user.User{},
// 		user.Country{},
// 		user.Province{},
// 		user.City{},
// 		user.District{},
// 	)
// }
