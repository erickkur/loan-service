package postgres

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/loan-service/internal/env"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Database struct {
	Conn *bun.DB
	Tx   *bun.Tx
}

func NewDatabase() *Database {
	db := &Database{}

	return db
}

// HealthCheck ...
func (db *Database) HealthCheck() error {
	_, err := db.Conn.Exec("SELECT 1")
	if err != nil {
		fmt.Println("PostgreSQL is down")
	}

	return nil
}

// Connect ...
func (db *Database) Connect() {
	var tlsConfig *tls.Config

	if env.Env() == "production" || env.Env() == "staging" {
		tlsConfig = &tls.Config{InsecureSkipVerify: true}
	}

	parsedURL, err := url.Parse(env.DBUrl())
	if err != nil {
		log.Fatalf("Error parsing database URL: %v", err)
	}

	connector := pgdriver.NewConnector(
		pgdriver.WithApplicationName(env.AppName()),
		pgdriver.WithDSN(parsedURL.String()),
		pgdriver.WithTLSConfig(tlsConfig),
	)

	db.Conn = bun.NewDB(sql.OpenDB(connector), pgdialect.New())
}

// Close ...
func (db *Database) Close() {
	db.Conn.Close()
}
