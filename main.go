package main

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/article"
	"banco/domain/diskon"
	"banco/domain/inventori"
	"banco/domain/produk"
	"banco/domain/transaksi"
	"banco/domain/user"
	"banco/httproute/route"
	"banco/httproute/service"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	conf := config.GetConfig()

	r := gin.Default()
	s := service.NewService(conf)
	route.NewRoute(s, r)

	r.Run()
	createTable()

}

func createTable() {
	db, err := orm.DbCon(config.GetConfig())
	if err != nil {
		log.Println(err)
		return
	}

	db.AutoMigrate(
		article.Article{},
		diskon.Discount{},
		inventori.Inventory{},
		inventori.HistoryOutcomingItem{},
		produk.Product{},
		produk.ProductDiscount{},
		produk.ProductType{},
		transaksi.Transaction{},
		transaksi.TransactionDetail{},
		transaksi.Amount{},
		transaksi.Satuan{},
		user.User{},
		user.Country{},
		user.Province{},
		user.City{},
		user.District{},
	)

	// person := []user.User{{
	// 	Model:      orm.Model{},
	// 	Name:       "Agung",
	// 	Email:      "agung@gmail.com",
	// 	Password:   HashPassword(),
	// 	CountryId:  0,
	// 	ProvinceId: 0,
	// 	CityId:     0,
	// 	DistrictId: 0,
	// 	Address:    "jalan kemayoran no 10",
	// }, {
	// 	Model:      orm.Model{},
	// 	Name:       "Nurcahyo Rosiandana",
	// 	Email:      "nurcahyo@gmail.com",
	// 	Password:   HashPassword(),
	// 	CountryId:  0,
	// 	ProvinceId: 0,
	// 	CityId:     0,
	// 	DistrictId: 0,
	// 	Address:    "",
	// }}
	// db.Create(&person)

	// article := []article.Article{
	// 	{
	// 		Model: orm.Model{
	// 			CraetedBy: 1,
	// 		},
	// 		Title:       "Ini merupakan sebuah article pertama",
	// 		Description: "Deskripsi ini adalah lanjutan dari sebuah title",
	// 		TimeStart:   time.Now(),
	// 		TimeEnd:     time.Now().AddDate(0, 5, 0),
	// 	}, {
	// 		Model: orm.Model{
	// 			CraetedBy: 1,
	// 		},
	// 		Title:       "article kedua dari sebuah judul",
	// 		Description: "deskripsi dari sebuah deskripsi",
	// 		TimeStart:   time.Now().AddDate(0, 1, 0),
	// 		TimeEnd:     time.Now().AddDate(0, 5, 0),
	// 	}}
	// db.Create(&article)
}

func HashPassword() string {
	password := "password"
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return "password"
	}

	return string(result)
}

func CompareHashPassword(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
