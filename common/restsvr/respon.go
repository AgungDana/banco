package restsvr

import (
	"banco/common/werror"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	SUCCES = "success"
	FAILED = "failed"
)

type ErrorResponse struct {
	Code    string
	Message string
	Detail  map[string]string
}

type Repsonse struct {
	Status string
	Data   interface{}
	Errors []ErrorResponse
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
	if data == nil {
		r.Status = FAILED
	}
	if err != nil {
		r.generatedError(err)
	}

	r.Data = data
}

func CreateResponse(c *gin.Context, res *Repsonse) {
	c.JSON(200, res)
}
