package loan

import (
	model "github.com/loan-service/adapter/models/loan"
	"github.com/loan-service/application/dto"
)

type LoanServiceInterface interface {
	CreateLoan(request dto.CreateLoanRequest) (*model.Loan, error)
}
