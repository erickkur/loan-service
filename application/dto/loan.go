package dto

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/loan-service/internal/constant"
	errs "github.com/loan-service/internal/error"
)

type CreateLoanRequest struct {
	BorrowerGUID       uuid.UUID `json:"borrowerGUID"`
	PrincipalAmount    int64     `json:"principalAmount"`
	Rate               float64   `json:"rate"`
	ReturnOfInvestment int64     `json:"returnOfInvestment"`
}

type CreateLoanResponse struct {
	GUID uuid.UUID `json:"guid"`
}

type UpdateLoanResponse struct {
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
	LoanGUID           *uuid.UUID `json:"loanGUID"`
	Status             string     `json:"status"`
	PictureProof       string     `json:"pictureProof"`
	EmployeeGUID       uuid.UUID  `json:"employeeGUID"`
	DateOfApproval     time.Time  `json:"dateOfApproval"`
	DateOfDisbursement time.Time  `json:"dateOfDisbursement"`
	AgreementLetter    string     `json:"agreementLetter"`
}

func isLoanStatusExist(status string) bool {
	return status == constant.LoanStatusApproved || status == constant.LoanStatusInvested || status == constant.LoanStatusDisbursed
}

func checkLoanStatusCanBeProcessToApproved(r UpdateLoanRequest) []string {
	var invalidFields []string

	if r.Status != constant.LoanStatusApproved {
		return invalidFields
	}

	if r.PictureProof == "" {
		invalidFields = append(invalidFields, "pictureProof")
	}

	if !isImageFile(r.PictureProof) {
		invalidFields = append(invalidFields, "agreementLetter extension must be image")
	}

	if r.EmployeeGUID == uuid.Nil {
		invalidFields = append(invalidFields, "employeeGUID")
	}

	if r.DateOfApproval.IsZero() {
		invalidFields = append(invalidFields, "dateOfApproval")
	}

	return invalidFields
}

func checkLoanStatusCanBeProcessToDisbursed(r UpdateLoanRequest) []string {
	var invalidFields []string

	if r.Status != constant.LoanStatusDisbursed {
		return invalidFields
	}

	if r.AgreementLetter == "" {
		invalidFields = append(invalidFields, "agreementLetter")
	}

	if !isPDFFile(r.AgreementLetter) {
		invalidFields = append(invalidFields, "agreementLetter extension must be .pdf")
	}

	if r.EmployeeGUID == uuid.Nil {
		invalidFields = append(invalidFields, "employeeGUID")
	}

	if r.DateOfDisbursement.IsZero() {
		invalidFields = append(invalidFields, "dateOfDisbursement")
	}

	return invalidFields
}

func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg":
		return true
	default:
		return false
	}
}

func isPDFFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))

	if ext == ".pdf" {
		return true
	}

	return false
}

func (r UpdateLoanRequest) Validate() error {
	var invalidFields []string
	if !isLoanStatusExist(r.Status) {
		return errs.ValidationAcceptedValue{Field: "status"}
	}

	invalidFields = checkLoanStatusCanBeProcessToApproved(r)
	if len(invalidFields) > 0 {
		return errs.ValidationRequiredData{InvalidFields: invalidFields}
	}

	invalidFields = checkLoanStatusCanBeProcessToDisbursed(r)
	if len(invalidFields) > 0 {
		return errs.ValidationRequiredData{InvalidFields: invalidFields}
	}

	return nil
}
