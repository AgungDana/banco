package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/common/werror"
	"banco/domain/article"
	"errors"
	"strconv"
	"time"

	"gorm.io/gorm"
)

//no mutex
func NewRepoArticle(conf config.Config) (article.ArticleRepo, error) {
	db, err := orm.DbCon(conf)
	if err != nil {
		return nil, err
	}
	return &repo{
		db: db,
	}, nil
}

type repo struct {
	db *gorm.DB
}

// Create implements artikel.ArticleRepo
func (r *repo) Create(in article.Article) (*uint, error) {
	//create article from input
	err := r.db.Model(article.Article{}).Create(in).Error
	if err != nil {
		return nil, err
	}
	return &in.Id, nil
}

// Delete implements artikel.ArticleRepo
func (r *repo) DeleteById(i uint) (*uint, error) {

	err := r.db.Delete(&article.Article{}, i).Error
	if err != nil {
		err = werror.Error{ //custom error
			Code: "INTERNAL SERVER ERROR", Message: "error delete article", Detail: map[string]string{"error": err.Error(), "id": strconv.Itoa(int(i))},
		}
		return nil, err
	}
	return &i, nil
}

// Read implements artikel.ArticleRepo
func (r *repo) FindById(i uint) (*article.Article, error) {
	data := article.Article{}
	err := r.db.Where("id = ? AND time_start >= ? AND time_end <= ?", i, time.Now(), time.Now()).Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &data, err

}

// Read implements artikel.ArticleRepo
func (r *repo) FindAll() ([]*article.Article, error) {
	data := []*article.Article{}
	err := r.db.Where("time_start >= ? AND time_end <= ?", time.Now(), time.Now()).Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("not found")
	}
	return data, nil
}

// Update implements artikel.ArticleRepo
func (r *repo) Update(in article.Article) (*uint, error) {
	err := r.db.Model(article.Article{}).Updates(in).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return &in.Id, nil
}
