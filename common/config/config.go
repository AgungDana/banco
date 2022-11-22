package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var conf Config

type Config struct {
	Host      string
	Port      string
	DbName    string
	UserName  string
	Password  string
	Driver    string
	Address   string
	ImagePath string
	ImageUrl  string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")
	userName := os.Getenv("DBUSERNAME")
	password := os.Getenv("DBPASSWORD")
	driver := os.Getenv("DBDRIVER")
	address := os.Getenv("ADDRESS")
	imagePath := os.Getenv("IMAGEPATH")
	imageUrl := os.Getenv("IMAGEURL")
	conf = Config{
		Host:      host,
		Port:      port,
		DbName:    dbName,
		UserName:  userName,
		Password:  password,
		Driver:    driver,
		Address:   address,
		ImagePath: imagePath,
		ImageUrl:  imageUrl,
	}

}

func GetConfig() Config {
	return conf
}
