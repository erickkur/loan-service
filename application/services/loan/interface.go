package loan

import (
	"context"

	"github.com/google/uuid"
	loanAdapter "github.com/loan-service/adapter/models/loan"
	"github.com/loan-service/application/dto"
)

type LoanServiceInterface interface {
	CreateLoan(ctx context.Context, request dto.CreateLoanRequest) (*loanAdapter.Loan, error)
	UpdateLoanToApprovedOrDisbursed(ctx context.Context, request dto.UpdateLoanRequest) (*loanAdapter.Loan, error)
	GetLoanByGUID(ctx context.Context, loanGUID uuid.UUID) (*loanAdapter.Loan, error)
}
