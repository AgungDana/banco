package article

import "banco/common/infra/orm"

type Artikel struct {
	orm.Model
	Title     string
	Deskripsi string
}

func (*Artikel) TableName() string {
	return "artikel"
}

type ArticleResponse struct {
	Title     string
	Deskripsi string
}
