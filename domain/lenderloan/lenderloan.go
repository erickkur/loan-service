package lenderloan

import (
	"context"
	"time"

	"github.com/loan-service/application/dto"
	llService "github.com/loan-service/application/services/lenderloan"
	lService "github.com/loan-service/application/services/loan"
	llogService "github.com/loan-service/application/services/loanlog"
	"github.com/loan-service/internal/constant"
	errs "github.com/loan-service/internal/error"
)

type EntityDependency struct {
	LenderLoanService llService.LenderLoanServiceInterface
	LoanService       lService.LoanServiceInterface
	LoanLogService    llogService.LoanLogServiceInterface
}

type Entity struct {
	lenderLoanService llService.LenderLoanServiceInterface
	loanService       lService.LoanServiceInterface
	loanLogService    llogService.LoanLogServiceInterface
}

func NewEntity(d EntityDependency) Entity {
	return Entity{
		lenderLoanService: d.LenderLoanService,
		loanService:       d.LoanService,
		loanLogService:    d.LoanLogService,
	}
}

func (l Entity) CreateLenderLoan(ctx context.Context, request dto.CreateLenderLoanRequest) (*dto.CreateLenderLoanResponse, *errs.JSONWrapError) {
	var response dto.CreateLenderLoanResponse

	err := request.Validate()
	if err != nil {
		jsonWrapErr := errs.NewValidationError(err).WrapError(errs.LoanPrefix)
		return nil, &jsonWrapErr
	}

	loan, err := l.loanService.GetLoanByGUID(ctx, request.LoanGUID)
	if err != nil {
		jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanPrefix)
		return nil, &jsonWrapErr
	}

	latestLog, err := l.loanLogService.GetLatestLoanLog(ctx, int64(loan.ID))
	if err != nil {
		jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanLogPrefix)
		return nil, &jsonWrapErr
	}

	if latestLog.Status != constant.LoanStatusApproved {
		err := errs.CustomErrorInformation{
			ErrorInformation: "current loan status must be approved",
		}
		jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanLogPrefix)
		return nil, &jsonWrapErr
	}

	lenderLoans, err := l.lenderLoanService.GetLenderLoansByLoanID(ctx, int64(loan.ID))
	if err != nil {
		jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LenderLoanPrefix)
		return nil, &jsonWrapErr
	}

	var currentInvestedAmount int64
	for _, lenderLoan := range lenderLoans {
		currentInvestedAmount += lenderLoan.InvestedAmount
	}

	totalInvestedAmount := currentInvestedAmount + request.InvestedAmount
	if totalInvestedAmount <= loan.PrincipalAmount {
		lenderLoan, err := l.lenderLoanService.CreateLenderLoan(ctx, request, loan.ID)
		if err != nil {
			jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LenderLoanPrefix)
			return nil, &jsonWrapErr
		}

		response = dto.CreateLenderLoanResponse{
			GUID: lenderLoan.GUID,
		}
	}

	if totalInvestedAmount == loan.PrincipalAmount {
		_, err := l.loanLogService.CreateLoanLog(ctx, dto.CreateLoanLogParam{
			LoanID:    loan.ID,
			Status:    constant.LoanStatusInvested,
			CreatedAt: time.Now(),
		})
		if err != nil {
			jsonWrapErr := errs.NewDatabaseError(err).WrapError(errs.LoanLogPrefix)
			return nil, &jsonWrapErr
		}
	}

	return &response, nil
}
