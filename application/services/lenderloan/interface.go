package lenderloan

import (
	"context"

	lenderLoanAdapter "github.com/loan-service/adapter/models/lenderloan"
	"github.com/loan-service/application/dto"
)

type LenderLoanServiceInterface interface {
	CreateLenderLoan(ctx context.Context, request dto.CreateLenderLoanRequest, loanID int) (*lenderLoanAdapter.LenderLoan, error)
	GetLenderLoansByLoanID(ctx context.Context, loanID int64) ([]lenderLoanAdapter.LenderLoan, error)
}
