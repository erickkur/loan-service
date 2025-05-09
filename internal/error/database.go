package error

import (
	"database/sql"
	"net/http"

	"github.com/uptrace/bun/driver/pgdriver"
)

type DatabaseError struct {
	error
}

func NewDatabaseError(err error) DatabaseError {
	return DatabaseError{err}
}

func (d DatabaseError) WrapError(domainCode string) JSONWrapError {
	pgErr, ok := d.error.(pgdriver.Error)
	if ok && pgErr.IntegrityViolation() {
		return JSONWrapError{
			Error:   d.error,
			Status:  http.StatusInternalServerError,
			Code:    generateErrorCode(domainCode, PgIntegrityViolation),
			Message: d.Error(),
		}
	}

	if d.error == sql.ErrNoRows {
		return JSONWrapError{
			Error:   d.error,
			Status:  http.StatusNotFound,
			Code:    generateErrorCode(domainCode, PgErrNoRows),
			Message: "Entry not found",
		}
	}

	_, isCustomErrorInformation := d.error.(CustomErrorInformation)
	if isCustomErrorInformation {
		return JSONWrapError{
			Error:   d.error,
			Status:  http.StatusBadRequest,
			Code:    generateErrorCode(domainCode, UnacceptedValueError),
			Message: d.Error(),
		}
	}

	return JSONWrapError{
		Error:   d.error,
		Status:  http.StatusInternalServerError,
		Code:    generateErrorCode(domainCode, PgUnknownError),
		Message: "Unknown database error",
	}
}
