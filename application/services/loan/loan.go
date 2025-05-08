package loan

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	model "github.com/loan-service/adapter/models/loan"
	"github.com/loan-service/application/dto"
)

type Dependency struct {
	LoanModel model.LoanModelInterface
	DBClient  pg.DatabaseAdapterInterface
}

type LoanService struct {
	loanModel model.LoanModelInterface
	dbClient  pg.DatabaseAdapterInterface
}

func NewLoanService(d Dependency) *LoanService {
	return &LoanService{
		loanModel: d.LoanModel,
		dbClient:  d.DBClient,
	}
}

func (ls *LoanService) CreateLoan(request dto.CreateLoanRequest) (*model.Loan, error) {
	l := model.Loan{
		BorrowerID:         request.BorrowerID,
		PrincipalAmount:    request.PrincipalAmount,
		Rate:               request.Rate,
		ReturnOfInvestment: request.ReturnOfInvestment,
	}

	return ls.loanModel.CreateLoan(
		ls.dbClient,
		context.TODO(),
		l,
	)
}
