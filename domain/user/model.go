package user

import (
	"banco/common/infra/orm"
)

type User struct {
	orm.Model
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	CountryId  int    `json:"countryId,omitempty"`
	ProvinceId int    `json:"provinceId,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	DistrictId int    `json:"districtId,omitempty"`
	Address    string `json:"address,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type Profile struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	CountryId  int    `json:"countryId,omitempty"`
	ProvinceId int    `json:"provinceId,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	DistrictId int    `json:"districtId,omitempty"`
	Address    string `json:"address,omitempty"`
}

func ToProfile(data User) Profile {
	return Profile{
		Name:       data.Name,
		Email:      data.Email,
		CountryId:  data.CountryId,
		ProvinceId: data.ProvinceId,
		CityId:     data.CityId,
		DistrictId: data.DistrictId,
		Address:    data.Address,
	}
}

type CreateUserRequest struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	CountryId  int    `json:"countryId,omitempty"`
	ProvinceId int    `json:"provinceId,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	DistrictId int    `json:"districtId,omitempty"`
	Address    string `json:"address,omitempty"`
}

func NewUser(req CreateUserRequest) User {
	return User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		CountryId:  req.CountryId,
		ProvinceId: req.ProvinceId,
		CityId:     req.CityId,
		DistrictId: req.DistrictId,
		Address:    req.Address,
	}
}

type UpdateUserRequest struct {
	Name       string `json:"name,omitempty"`
	CountryId  int    `json:"countryId,omitempty"`
	ProvinceId int    `json:"provinceId,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	DistrictId int    `json:"districtId,omitempty"`
	Address    string `json:"address,omitempty"`
}

func UpdateUser(req UpdateUserRequest, id uint) User {
	return User{
		Model: orm.Model{
			Id:        id,
			UpdatedBy: id,
		},
		Name:       req.Name,
		CountryId:  req.CountryId,
		ProvinceId: req.ProvinceId,
		CityId:     req.CityId,
		DistrictId: req.DistrictId,
		Address:    req.Address,
	}
}

type Country struct {
	orm.Model
	Name string `json:"name,omitempty"`
}

type Province struct {
	orm.Model
	CountryId int    `json:"countryId,omitempty"`
	Name      string `json:"name,omitempty"`
}

type City struct {
	orm.Model
	ProvinceId int    `json:"provinceId,omitempty"`
	Name       string `json:"name,omitempty"`
}

type District struct {
	orm.Model
	CityId int    `json:"cityId,omitempty"`
	Name   string `json:"name,omitempty"`
}
