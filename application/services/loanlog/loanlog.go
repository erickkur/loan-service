package loanlog

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	loanlogAdapter "github.com/loan-service/adapter/models/loanlog"
	"github.com/loan-service/application/dto"
)

type Dependency struct {
	LoanlogModel loanlogAdapter.LoanLogModelInterface
	DBClient     pg.DatabaseAdapterInterface
}

type LoanLogService struct {
	loanlogModel loanlogAdapter.LoanLogModelInterface
	dbClient     pg.DatabaseAdapterInterface
}

func NewLoanLogService(d Dependency) *LoanLogService {
	return &LoanLogService{
		loanlogModel: d.LoanlogModel,
		dbClient:     d.DBClient,
	}
}

func (ll *LoanLogService) GetLatestLoanLog(ctx context.Context, loanID int64) (*loanlogAdapter.LoanLog, error) {
	return ll.loanlogModel.GetLatestLoanLog(ll.dbClient, ctx, loanID)
}

func (ll *LoanLogService) CreateLoanLog(ctx context.Context, param dto.CreateLoanLogParam) (*loanlogAdapter.LoanLog, error) {
	lg := loanlogAdapter.LoanLog{
		LoanID:     int64(param.LoanID),
		Status:     param.Status,
		EmployeeID: param.EmployeeID,
		CreatedAt:  param.CreatedAt,
	}

	return ll.loanlogModel.CreateLoanLog(ll.dbClient, ctx, lg)
}
