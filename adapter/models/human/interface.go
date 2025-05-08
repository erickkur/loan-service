package human

import pg "github.com/loan-service/adapter/database/postgres"

type HumanModelInterface interface {
	GetHumansData(dbClient pg.DatabaseAdapterInterface, limit int) ([]Human, error)
}
