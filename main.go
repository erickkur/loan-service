package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/loan-service/adapter/httpserver"
	"github.com/loan-service/adapter/middleware"
	"github.com/loan-service/infra"
	"github.com/loan-service/internal/env"
	"github.com/loan-service/internal/logger"

	pg "github.com/loan-service/adapter/database/postgres"
	bwModel "github.com/loan-service/adapter/models/borrower"
	eModel "github.com/loan-service/adapter/models/employee"
	lModel "github.com/loan-service/adapter/models/loan"
	llogModel "github.com/loan-service/adapter/models/loanlog"
	lService "github.com/loan-service/application/services/loan"
	rs "github.com/loan-service/application/services/router"
	lDomain "github.com/loan-service/domain/loan"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	addr := flag.String("addr", env.AppPort(), "http service address")

	// Internal shared package
	log := logger.New()

	// Infra layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	infraObj := infra.Init()
	infraObj.Database.Connect()
	// ++++++++++++++++++++++++++++++++++++++++++

	// Adapter layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	httpServerAdapter := httpserver.NewAdapter(&httpserver.Adapter{
		Router: infraObj.Router,
		Log:    log,
	})
	middlewareAdapter := middleware.NewAdapter()
	postgresAdapter := pg.NewAdapter(infraObj.Database)

	loanModel := lModel.NewModel()
	borrowerModel := bwModel.NewModel()
	loanlog := llogModel.NewModel()
	employeeModel := eModel.NewModel()
	// ++++++++++++++++++++++++++++++++++++++++++

	// Service layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	routerService := rs.NewService(
		httpServerAdapter.Server,
		httpserver.RouterHelper{},
		middlewareAdapter,
		"/api")

	loanService := lService.NewLoanService(lService.Dependency{
		LoanModel:     loanModel,
		BorrowerModel: borrowerModel,
		LoanLogModel:  loanlog,
		EmployeeModel: employeeModel,
		DBClient:      postgresAdapter,
	})
	// ++++++++++++++++++++++++++++++++++++++++++

	// Domain layer initialization
	// ++++++++++++++++++++++++++++++++++++++++++
	lDomain.NewDomain(lDomain.RouteDependency{
		LoanService: loanService,
		Context:     routerService,
		Logger:      log,
	})
	// ++++++++++++++++++++++++++++++++++++++++++

	infraObj.Router.ReArrange()

	s := &http.Server{
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         *addr,
		Handler:      infraObj.Router,
	}

	log.Info(fmt.Sprint("Loan service started on port", env.AppPort()))

	return s.ListenAndServe()
}
