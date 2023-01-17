package main

import (
	"banco/common/config"
	"banco/httproute/route"
	"banco/httproute/service"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.GetConfig()
	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(CORSMiddleware())
	s := service.NewService(conf)
	route.NewRoute(conf, s, r)
	//
	r.Run(conf.Address)
	// createTable()
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
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
// produk.Product{},
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
