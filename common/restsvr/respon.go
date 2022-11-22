package restsvr

import "banco/common/werror"

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

func (r *Repsonse) ErrorFromWerror(errs error) ErrorResponse {
	err := ErrorResponse{}
	switch e := errs.(type) {
	case werror.Error:
		err = ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
			Detail:  e.Detail,
		}
	}
	return err
}
