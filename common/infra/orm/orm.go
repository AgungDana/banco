package orm

import (
	"banco/common/config"
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbCon(conf config.Config) (*gorm.DB, error) {
	var dsn string
	switch conf.Driver {
	case "mysql":
		dsn = conf.UserName + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" + conf.DbName + "?parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, errors.New("invalid driver")
	}
}
