package borrower

import (
	"github.com/loan-service/adapter/models/recordtimestamp"
	"github.com/uptrace/bun"
)

type Borrower struct {
	ID        int    `bun:",pk,autoincrement"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	recordtimestamp.RecordTimestamp

	bun.BaseModel `bun:"borrowers,alias:b"`
}
