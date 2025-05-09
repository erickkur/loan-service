package dto

import (
	"database/sql"
	"time"
)

type CreateLoanLogParam struct {
	LoanID     int
	Status     string
	EmployeeID sql.NullInt64
	CreatedAt  time.Time
}
