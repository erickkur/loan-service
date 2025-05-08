dep: ## Get the dependencies
	@go get -v ./...

run:
	source local-env.sh && ./tools/run.sh

migrate:
	source local-env.sh && ./tools/postgres-migrate.sh

test:
	go test ./...