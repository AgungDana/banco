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

func NewRepoArticle(conf config.Config) article.ArticleRepo {
	db, err := orm.DbCon(conf)
	if err != nil {
		panic(err)
	}
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

// Create implements artikel.ArticleRepo
func (r *repo) Create(article.Article) (*uint, error) {
	panic("unimplemented")
}

// Delete implements artikel.ArticleRepo
func (r *repo) Delete(i uint) error {
	// err := r.db.Where("where id = ?", i).Delete(&article.Artikel{Model: orm.Model{Id: i}}).Error
	err := r.db.Delete(&article.Article{}, i).Error
	if err != nil {
		err = werror.Error{
			Code: "INTERNAL SERVER ERROR", Message: "error delete article", Detail: map[string]string{"error": err.Error(), "id": strconv.Itoa(int(i))},
		}
		return err
	}
	return nil
}

// Read implements artikel.ArticleRepo
func (r *repo) FindById(i uint) (*article.Article, error) {
	data := article.Article{}
	err := r.db.Where("id = ? AND time_start >= ? AND time_end =< ?", i, time.Now(), time.Now()).Find(&data).Error
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
	err := r.db.Find(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not found")
		}
		return nil, err
	}
	return data, err
}

// Update implements artikel.ArticleRepo
func (r *repo) Update(article.Article) (*uint, error) {
	panic("unimplemented")
}
