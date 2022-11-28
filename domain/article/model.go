package article

import (
	"banco/common/infra/orm"
	"time"
)

type Artikel struct {
	orm.Model
	Title       string
	Description string
	TimeStart   time.Time
	TimeEnd     time.Time
}

func (*Artikel) TableName() string {
	return "artikel"
}

type ArticleRequest struct {
	ArticleId uint `json:"articleId"`
}

type ArticleResponse struct {
	Id          uint
	Title       string
	Description string
	TimeStart   time.Time
	TimeEnd     time.Time
}
