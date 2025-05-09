package error

import (
	"fmt"
	"strconv"
	"strings"
)

// Unknown ...
const Unknown = 0

// ServiceID ...
const ServiceID = "1"

// Table prefix error code
const (
	LoanPrefix = "001"
)

// Own validation error code
const (
	UnknownValidationError = "500"
	ValidationDataError    = "501"
	UnacceptedValueError   = "502"
)

// General error code
const (
	GeneralUnknownError  = "000"
	PgIntegrityViolation = "100"
	PgErrNoRows          = "101"
	PgErrMultiRows       = "102"
	PgUnknownError       = "103"
	DecoderSyntaxError   = "104"
	DecoderUnknownError  = "105"
	DecoderMultiError    = "106"
)

// Error type
const (
	PgErrorType      = "pgErrorType"
	DecoderErrorType = "decoderErrorType"
	StrconvErrorType = "strconvErrorType"
)

var generalErrorMap = map[string]string{
	PgErrorType:      "pgErrorType",
	DecoderErrorType: "decoderErrorType",
	StrconvErrorType: "strconvErrorType",
}

type ValidationRequiredData struct {
	InvalidFields []string
}

type ValidationAcceptedValue struct {
	Field string
}

type CustomErrorInformation struct {
	ErrorInformation string
}

func (e ValidationRequiredData) Error() string {
	return fmt.Sprintf("The following fields is required: %s", strings.Join(e.InvalidFields[:], ","))
}

func (e ValidationAcceptedValue) Error() string {
	return fmt.Sprintf("Unaccepted value on param %s", e.Field)
}

func (e CustomErrorInformation) Error() string {
	return e.ErrorInformation
}

func generateErrorCode(domainCode string, suffixCode string) int {
	s := fmt.Sprintf("%s%s%s", ServiceID, domainCode, suffixCode)
	code, err := strconv.Atoi(s)
	if err != nil {
		return Unknown
	}

	return code
}
