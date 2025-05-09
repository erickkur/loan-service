package dto

import (
	"github.com/google/uuid"
	"github.com/loan-service/internal/constant"
	errs "github.com/loan-service/internal/error"
)

type CreateLoanRequest struct {
	BorrowerGUID       uuid.UUID `json:"borrowerGUID"`
	PrincipalAmount    float64   `json:"principalAmount"`
	Rate               float64   `json:"rate"`
	ReturnOfInvestment float64   `json:"returnOfInvestment"`
}

type CreateLoanResponse struct {
	GUID uuid.UUID `json:"guid"`
}

func (r CreateLoanRequest) Validate() error {
	var invalidFields []string
	if r.BorrowerGUID == uuid.Nil {
		invalidFields = append(invalidFields, "borrowerGUID")
	}

	if r.PrincipalAmount == 0 {
		invalidFields = append(invalidFields, "principalAmount")
	}

	if r.Rate == 0 {
		invalidFields = append(invalidFields, "rate")
	}

	if r.ReturnOfInvestment == 0 {
		invalidFields = append(invalidFields, "returnOfInvestment")
	}

	if len(invalidFields) > 0 {
		return errs.ValidationRequiredData{InvalidFields: invalidFields}
	}

	return nil
}

type UpdateLoanRequest struct {
	Status string `json:"status,omitempty"`
}

func isLoanStatusExist(status string) bool {
	return status != constant.LoanStatusApproved && status != constant.LoanStatusInvested && status != constant.LoanStatusDisbursed
}

func (r UpdateLoanRequest) Validate() error {
	var invalidFields []string
	if !isLoanStatusExist(r.Status) {
		invalidFields = append(invalidFields, "loan_status")
	}

	return nil
}
