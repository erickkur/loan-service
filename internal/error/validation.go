package error

import "net/http"

type ValidationError struct {
	error
}

func NewValidationError(error error) ValidationError {
	return ValidationError{error: error}
}

func (v ValidationError) WrapError(domainCode string) JSONWrapError {
	_, isValidationRequiredData := v.error.(ValidationRequiredData)
	_, isValidationAcceptedValue := v.error.(ValidationAcceptedValue)

	if isValidationRequiredData || isValidationAcceptedValue {
		return JSONWrapError{
			Error:   v.error,
			Status:  http.StatusBadRequest,
			Code:    generateErrorCode(domainCode, ValidationDataError),
			Message: v.Error(),
		}
	}

	return JSONWrapError{
		Error:   v.error,
		Status:  http.StatusInternalServerError,
		Code:    generateErrorCode(domainCode, UnknownValidationError),
		Message: v.Error(),
	}
}
