package postgres

type PostgresInfraInterface interface {
	Connect()
	Close()
	HealthCheck() error
}
