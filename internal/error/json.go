package error

import "fmt"

type JSONWrapError struct {
	Error   error
	Status  int
	Code    int
	Message string
}

func NewJsonWrapErrorService() JSONWrapError {
	return JSONWrapError{}
}

// StringWithError get full string with error code, message and error object
func (e JSONWrapError) StringWithError() string {
	return fmt.Sprintf("Error Code: %d, Message: %s. Detail: %s", e.Code, e.Message, e.Error.Error())
}
