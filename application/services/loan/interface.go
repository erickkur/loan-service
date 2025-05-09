package loan

import (
	"context"

	loanAdapter "github.com/loan-service/adapter/models/loan"
	"github.com/loan-service/application/dto"
)

type LoanServiceInterface interface {
	CreateLoan(ctx context.Context, request dto.CreateLoanRequest) (*loanAdapter.Loan, error)
	UpdateLoanToApproved(ctx context.Context, request dto.UpdateLoanRequest) (*loanAdapter.Loan, error)
}
