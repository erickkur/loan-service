package loanlog

import (
	"context"
	"database/sql"
	"fmt"
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
		fmt.Println(err)
		return nil, err
	}

	query := db.GetConnectionDB().
		NewInsert().
		Model(&lg)
	_, err = query.Exec(ctx)
	fmt.Println(err)

	return &lg, err
}
