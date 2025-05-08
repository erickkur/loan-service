package error

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
)

type DecoderError struct {
	error
}

func NewDecoderError(error error) DecoderError {
	return DecoderError{error: error}
}

func (d DecoderError) WrapError(domainCode string) JSONWrapError {
	if _, ok := d.error.(*json.SyntaxError); ok {
		return JSONWrapError{
			Error:   d.error,
			Status:  http.StatusBadRequest,
			Code:    generateErrorCode(domainCode, DecoderSyntaxError),
			Message: d.Error(),
		}
	}

	if _, ok := d.error.(schema.MultiError); ok {
		return JSONWrapError{
			Error:   d.error,
			Status:  http.StatusBadRequest,
			Code:    generateErrorCode(domainCode, DecoderMultiError),
			Message: d.Error(),
		}
	}

	return JSONWrapError{
		Error:   d.error,
		Status:  http.StatusInternalServerError,
		Code:    generateErrorCode(domainCode, DecoderUnknownError),
		Message: d.Error(),
	}
}
