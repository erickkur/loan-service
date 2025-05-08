package loan

import (
	"context"

	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Loan struct {
	ID                  int     `bun:",pk,autoincrement"`
	BorrowerID          int64   `json:"borrower_id"`
	PrincipalAmount     float64 `json:"principal_amount"`
	Rate                float64 `json:"rate"`
	ReturnOfInvestment  float64 `json:"return_of_investment"`
	AgreementLetterLink string  `json:"agreement_letter_link"`

	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"loans,alias:l"`
}

type LoanModel struct {
}

func NewModel() *LoanModel {
	return &LoanModel{}
}

func (m *LoanModel) CreateLoan(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	l Loan,
) (*Loan, error) {
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewInsert().
		Model(&l)
	_, err = query.Exec(ctx)

	return &l, err
}
