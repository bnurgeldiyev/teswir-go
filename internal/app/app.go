package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"os"
	"os/signal"
	"syscall"
	"teswir-go/config"
	v1 "teswir-go/internal/controller/http/v1"
	"teswir-go/internal/usecase"
	"teswir-go/internal/usecase/repo"
	"teswir-go/internal/usecase/webapi"
	"teswir-go/pkg/auth"
	"teswir-go/pkg/httpserver"
	"teswir-go/pkg/logger"
	"teswir-go/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	authGrpc, err1 := auth.NewGrpcClient(cfg.AUTH.URL, l)
	if err1 != nil {
		l.Fatal(fmt.Errorf("app - Run - auth.NewGrpcClient"))
	}

	userUseCase := usecase.NewUseCase(
		repo.NewRepo(pg),
		webapi.NewWebAPI(authGrpc),
	)

	handler := mux.NewRouter()
	v1.NewRouter(handler, l, userUseCase)

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	fmt.Println("<--START-SERVER-->")
	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	/*err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}*/
}
