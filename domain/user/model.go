package user

import "banco/common/infra/orm"

type User struct {
	orm.Model
	Name     string
	Email    string
	Password string
	Address  string
}

type Country struct {
	orm.Model
	Name string
}

type Province struct {
	orm.Model
	CountryId int
	Name      string
}

type City struct {
	orm.Model
	ProvinceId int
	Name       string
}

type District struct {
	orm.Model
	CityId int
	Name   string
}
