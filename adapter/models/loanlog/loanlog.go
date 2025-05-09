package loanlog

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	pg "github.com/loan-service/adapter/database/postgres"
	"github.com/uptrace/bun"
)

type LoanLog struct {
	ID         int           `bun:",pk,autoincrement"`
	GUID       uuid.UUID     `bun:",nullzero"`
	LoanID     int64         `bun:",notnull"`
	Status     string        `bun:",notnull,default:proposed"`
	EmployeeID sql.NullInt64 `bun:""`
	CreatedAt  time.Time     `bun:",,default:now()"`

	bun.BaseModel `bun:"loan_logs,alias:lg"`
}

type LoanLogModel struct {
}

func NewModel() *LoanLogModel {
	return &LoanLogModel{}
}

func (m *LoanLogModel) CreateLoanLog(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	lg LoanLog,
) (*LoanLog, error) {
	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewInsert().
		Model(&lg)
	_, err = query.Exec(ctx)

	return &lg, err
}

func (m *LoanLogModel) GetLatestLoanLog(
	dbClient pg.DatabaseAdapterInterface,
	ctx context.Context,
	loanID int64,
) (*LoanLog, error) {
	var ll LoanLog

	db, err := dbClient.Get()
	if err != nil {
		return nil, err
	}

	query := db.GetConnectionDB().
		NewSelect().
		Model(&ll).
		Where("loan_id = ?", loanID).
		Order("id desc").
		Limit(1)
	err = query.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &ll, nil
}
