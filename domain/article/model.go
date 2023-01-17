package article

import (
	"banco/common/infra/orm"
	"time"
)

type Article struct {
	orm.Model
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	TimeStart   time.Time `json:"timeStart,omitempty"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
}

func (*Article) TableName() string {
	return "artikel"
}

type CreateArticleRequest struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	TimeStart   time.Time `json:"timeStart,omitempty"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
}

func (c *CreateArticleRequest) NewArticle(createdBy uint) Article {
	return Article{
		Model: orm.Model{
			CraetedBy: createdBy,
		},
		Title:       c.Title,
		Description: c.Description,
		TimeStart:   c.TimeStart,
		TimeEnd:     c.TimeEnd,
	}
}

type UpdateArticleRequest struct {
	Id          uint      `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	TimeStart   time.Time `json:"timeStart,omitempty"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
}

func (c *UpdateArticleRequest) ChangeArticle(updatedBy uint) Article {
	return Article{
		Model: orm.Model{
			Id:        c.Id,
			UpdatedBy: updatedBy,
		},
		Title:       c.Title,
		Description: c.Description,
		TimeStart:   c.TimeStart,
		TimeEnd:     c.TimeEnd,
	}
}

type ArticleRequest struct {
	ArticleId uint `json:"articleId"`
}

type ArticleResponse struct {
	Id          uint      `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	TimeStart   time.Time `json:"timeStart,omitempty"`
	TimeEnd     time.Time `json:"timeEnd,omitempty"`
}

func ToResponse(a *Article) *ArticleResponse {
	return &ArticleResponse{
		Id:          a.Id,
		Title:       a.Title,
		Description: a.Description,
		TimeStart:   a.TimeStart,
		TimeEnd:     a.TimeEnd,
	}
}
