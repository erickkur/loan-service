package database

import "github.com/uptrace/bun"

type DatabaseAdapterInterface interface {
	Conn() *bun.DB
	Get() (*DatabaseAdapter, error)
	HealthCheck() error
	Commit() error
	Rollback() error
	BeginTransaction() (DatabaseAdapterInterface, error)
	GetConnectionDB() bun.IDB
}
