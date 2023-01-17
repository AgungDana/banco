package repo

import (
	"banco/domain/user"
	"context"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func newQuery(db *gorm.DB) user.Query {
	return &query{db: db}
}

// FindCities implements user.Query
func (q *query) FindCities(ctx context.Context, id uint) ([]*user.City, error) {
	// panic("unimplemented")
	kota := []*user.City{}
	err := q.db.Where("province_id = ?", id).Find(&kota).Error
	if err != nil {
		return nil, err
	}
	return kota, nil
}

// FindDistricts implements user.Query
func (q *query) FindDistricts(ctx context.Context, id uint) ([]*user.District, error) {
	// panic("unimplemented")
	kecamatan := []*user.District{}
	err := q.db.Where("city_id = ?", id).Find(&kecamatan).Error
	if err != nil {
		return nil, err
	}
	return kecamatan, nil
}

// FindProvinces implements user.Query
func (q *query) FindProvinces(ctx context.Context, id uint) ([]*user.Province, error) {
	provinsi := []*user.Province{}
	err := q.db.Where("country_id = ?", id).Find(&provinsi).Error
	if err != nil {
		return nil, err
	}
	return provinsi, nil
}

// GetUserByEmail implements user.Query
func (q *query) FindUserByEmail(ctx context.Context, email string) (*user.User, error) {
	person := new(user.User)
	err := q.db.Where("email = ?", email).Find(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}

// GetUserById implements user.Query
func (q *query) FindUserByID(ctx context.Context, id uint) (*user.User, error) {
	person := new(user.User)
	err := q.db.Where("id = ?", id).Preload("Country").Preload("Province").Preload("City").Preload("District").Find(person).Error
	if err != nil {
		return nil, err
	}
	return person, nil
}
