package diskon

import (
	"banco/common/infra/orm"
	"time"
)

type Discount struct {
	orm.Model
	Deskripsi     string
	Type          string
	Discount      int
	StartDiscount time.Time
	EndDiscount   time.Time
}
