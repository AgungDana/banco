package restsvr

import (
	"banco/common/werror"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

var (
	SUCCES = "success"
	FAILED = "failed"
)

type ErrorResponse struct {
	Code    string            `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Detail  map[string]string `json:"detail,omitempty"`
}

type Repsonse struct {
	Status string          `json:"status,omitempty"`
	Data   interface{}     `json:"data,omitempty"`
	Errors []ErrorResponse `json:"errors,omitempty"`
}

func errFromWerror(err werror.Error) ErrorResponse {
	return ErrorResponse{
		Code:    err.Code,
		Message: err.Message,
		Detail:  err.Detail,
	}
}

func (r *Repsonse) generatedError(errs error) {
	err := []ErrorResponse{}
	switch e := errs.(type) {
	case werror.Error:
		err = append(err, errFromWerror(e))
	case werror.Errors:
		for _, v := range e.Errors {
			err = append(err, errFromWerror(v))
		}
	case error:
		err = append(err, ErrorResponse{
			Message: e.Error(),
		})
	default:
		fmt.Println("invalid error")
	}

	r.Errors = err
}

func (r *Repsonse) Add(data any, err error) {
	r.Status = SUCCES
	if data == nil || reflect.ValueOf(data).IsNil() {
		r.Status = FAILED
	}
	if err != nil {
		r.generatedError(err)
	}

	r.Data = data
}

func CreateResponse(c *gin.Context, res *Repsonse) {
	// c.JSON(200, res)
	c.JSON(200, res)
}
