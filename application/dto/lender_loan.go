package dto

import (
	"github.com/google/uuid"
	errs "github.com/loan-service/internal/error"
)

type CreateLenderLoanRequest struct {
	LoanGUID       uuid.UUID `json:"loanGUID"`
	LenderGUID     uuid.UUID `json:"lenderGUID"`
	InvestedAmount int64     `json:"investedAmount"`
}

func (r CreateLenderLoanRequest) Validate() error {
	var invalidFields []string
	if r.LoanGUID == uuid.Nil {
		invalidFields = append(invalidFields, "loanGUID")
	}

	if r.LenderGUID == uuid.Nil {
		invalidFields = append(invalidFields, "lenderGUID")
	}

	if r.InvestedAmount == 0 {
		invalidFields = append(invalidFields, "investedAmount")
	}

	if len(invalidFields) > 0 {
		return errs.ValidationRequiredData{InvalidFields: invalidFields}
	}

	return nil
}

type CreateLenderLoanResponse struct {
	GUID uuid.UUID `json:"guid"`
}
