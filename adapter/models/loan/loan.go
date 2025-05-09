package loan

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Loan struct {
	ID                  int            `bun:",pk,autoincrement"`
	GUID                uuid.UUID      `bun:",nullzero"`
	BorrowerID          int64          `bun:",notnull"`
	PrincipalAmount     int64          `bun:",notnull"`
	Rate                float64        `bun:",notnull"`
	ReturnOfInvestment  int64          `bun:",notnull"`
	AgreementLetterLink sql.NullString `bun:""`

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

func (b *LoanModel) GetLoanByGUID(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	guid uuid.UUID,
) (*Loan, error) {
	var l Loan
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&l).
		Where("guid = ?", guid)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (b *LoanModel) UpdateLoanAgrrementLetter(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	loanID int,
	aggrementLetter string,
) error {
	db, err := dbClient.Get()
	if err != nil {
		return err
	}

	query := db.GetConnectionDB().
		NewUpdate().
		Model(&Loan{}).
		Where("id = ?", loanID).
		Set("agreement_letter_link = ?", aggrementLetter)
	_, err = query.Exec(ctx)

	return err
}
