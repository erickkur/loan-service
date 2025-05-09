package loanlog

import (
	"context"

	loanlogAdapter "github.com/loan-service/adapter/models/loanlog"
	"github.com/loan-service/application/dto"
)

type LoanLogServiceInterface interface {
	GetLatestLoanLog(ctx context.Context, loanID int64) (*loanlogAdapter.LoanLog, error)
	CreateLoanLog(ctx context.Context, param dto.CreateLoanLogParam) (*loanlogAdapter.LoanLog, error)
}
