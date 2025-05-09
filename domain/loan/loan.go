package loan

import (
	"context"

	"github.com/loan-service/application/dto"
	lService "github.com/loan-service/application/services/loan"
	"github.com/loan-service/internal/constant"
	errs "github.com/loan-service/internal/error"
)

type EntityDependency struct {
	LoanService lService.LoanServiceInterface
}

type Entity struct {
	loanService lService.LoanServiceInterface
}

func NewEntity(d EntityDependency) Entity {
	return Entity{
		loanService: d.LoanService,
	}
}

func (l Entity) CreateLoan(ctx context.Context, request dto.CreateLoanRequest) (*dto.CreateLoanResponse, *errs.JSONWrapError) {
	err := request.Validate()
	if err != nil {
		jsonWrapErr := errs.NewValidationError(err).WrapError(errs.LoanPrefix)
		return nil, &jsonWrapErr
	}

	loan, err := l.loanService.CreateLoan(ctx, request)
	if err != nil {
		jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanPrefix)
		return nil, &jsonWrapErr
	}

	return &dto.CreateLoanResponse{
		GUID: loan.GUID,
	}, nil
}

func (l Entity) UpdateLoan(ctx context.Context, request dto.UpdateLoanRequest) (*dto.UpdateLoanResponse, *errs.JSONWrapError) {
	var response dto.UpdateLoanResponse

	err := request.Validate()
	if err != nil {
		jsonWrapErr := errs.NewValidationError(err).WrapError(errs.LoanPrefix)
		return nil, &jsonWrapErr
	}

	if request.Status == constant.LoanStatusApproved {
		loan, err := l.loanService.UpdateLoanToApproved(ctx, request)
		if err != nil {
			jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanPrefix)
			return nil, &jsonWrapErr
		}

		response = dto.UpdateLoanResponse{
			GUID: loan.GUID,
		}
	}

	return &response, nil
}
