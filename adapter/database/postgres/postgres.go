package database

import (
	"fmt"

	"github.com/loan-service/infra/postgres"
	"github.com/uptrace/bun"
)

type DatabaseAdapter struct {
	Pg *postgres.Database
	Tx *bun.Tx
}

func NewAdapter(pg *postgres.Database) *DatabaseAdapter {
	return &DatabaseAdapter{
		Pg: pg,
	}
}

func (db *DatabaseAdapter) Conn() *bun.DB {
	return db.Pg.Conn
}

func (db *DatabaseAdapter) Get() (*DatabaseAdapter, error) {
	if db.Conn() != nil {
		return db, nil
	}

	return nil, fmt.Errorf("infra-postgress: connection unavailable")
}

func (db *DatabaseAdapter) HealthCheck() error {
	_, err := db.Conn().Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
		return err
	}

	return nil
}

func (db *DatabaseAdapter) BeginTransaction() (DatabaseAdapterInterface, error) {
	newAdapter := NewAdapter(db.Pg)

	tx, err := newAdapter.Conn().Begin()
	if err != nil {
		return nil, err
	}

	newAdapter.Tx = &tx

	return newAdapter, nil
}

func (db *DatabaseAdapter) Commit() error {
	if db.Tx == nil {
		return fmt.Errorf("transaction not ready")
	}

	err := db.Tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseAdapter) Rollback() error {
	if db.Tx == nil {
		return fmt.Errorf("transaction not ready")
	}

	err := db.Tx.Rollback()
	if err != nil {
		return err
	}

	return nil
}

func (db *DatabaseAdapter) GetConnectionDB() bun.IDB {
	if db.Tx != nil {
		return db.Tx
	}

	return db.Conn()
}
