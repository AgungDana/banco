package werror

import (
	"fmt"
	"log"
)

type Error struct {
	Code    string
	Message string
	Detail  map[string]string
}

func (e Error) Error() string {
	if e.Code == "" {
		return e.Message
	}
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

type Errors struct {
	Errors []Error
}

func (e Errors) Error() string {
	result := "["
	errLength := len(e.Errors)
	for i, v := range e.Errors {
		separator := ""
		if i < errLength-1 {
			separator = ", "
		}
		result += v.Error() + separator
	}
	result += "]"
	return result
}

func New() *Errors {
	return &Errors{make([]Error, 0)}
}

func (e *Errors) AddError(errs error) {
	switch err := errs.(type) {
	case Error:
		e.Errors = append(e.Errors, err)
	case *Error:
		e.Errors = append(e.Errors, *err)
	case Errors:
		e.Errors = append(e.Errors, err.Errors...)
	case *Errors:
		e.Errors = append(e.Errors, err.Errors...)
	case error:
		e.Errors = append(e.Errors, Error{
			Message: err.Error(),
		})
	default:
		log.Fatal("invalid err type")
	}
}

func (e *Errors) Return() error {
	if len(e.Errors) == 0 {
		return nil
	}
	return e
}
