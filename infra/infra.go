package infra

import (
	"github.com/loan-service/infra/postgres"
	"github.com/loan-service/infra/way"
)

type Infra struct {
	Database *postgres.Database
	Router   *way.Router
}

func Init() *Infra {
	database := postgres.NewDatabase()

	router := way.NewRouter()

	return &Infra{
		Database: database,
		Router:   router,
	}
}
