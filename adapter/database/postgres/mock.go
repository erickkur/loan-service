package database

import (
	"errors"

	"github.com/uptrace/bun"
)

type MockDatabaseAdapter struct {
	MockConn           *bun.DB
	MockTx             bun.Tx
	MockError          error
	MockBeginTxAdapter DatabaseAdapterInterface
}

func (m *MockDatabaseAdapter) Conn() *bun.DB {
	return m.MockConn
}

func (m *MockDatabaseAdapter) Get() (*DatabaseAdapter, error) {
	if m.MockConn != nil {
		return &DatabaseAdapter{}, nil
	}
	return nil, errors.New("mocked connection unavailable")
}

func (m *MockDatabaseAdapter) HealthCheck() error {
	return m.MockError
}

func (m *MockDatabaseAdapter) BeginTransaction() (DatabaseAdapterInterface, error) {
	if m.MockError != nil {
		return nil, m.MockError
	}
	return m.MockBeginTxAdapter, nil
}

func (m *MockDatabaseAdapter) Commit() error {
	return m.MockError
}

func (m *MockDatabaseAdapter) Rollback() error {
	return m.MockError
}

func (m *MockDatabaseAdapter) GetConnectionDB() bun.IDB {
	// if m.MockTx != nil {
	// 	return &m.MockTx
	// }
	return m.MockConn
}
