package app

import (
	"context"
	"net/http"

	"github.com/cucumberjaye/gophermart/configs"
	"github.com/cucumberjaye/gophermart/internal/app/handler"
	"github.com/cucumberjaye/gophermart/internal/app/repository/postgresdb"
	service "github.com/cucumberjaye/gophermart/internal/app/sevice"
	"github.com/cucumberjaye/gophermart/internal/app/worker"
	"github.com/cucumberjaye/gophermart/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type App struct {
	mux *chi.Mux
}

func New() (*App, error) {
	configs.InitConfigs()

	db, err := postgres.New()
	if err != nil {
		return nil, err
	}

	martRepository, err := postgresdb.New(db)
	if err != nil {
		return nil, err
	}

	martService := service.New(martRepository)

	martHandler := handler.New(martService)

	worker := worker.New(martRepository)
	go worker.Start(context.Background())

	router := chi.NewRouter()
	router.Mount("/api", martHandler.InitRoutes())

	return &App{mux: router}, nil
}

func (a *App) Run() error {
	log.Print("server running")

	return http.ListenAndServe(configs.RunAddress, a.mux)
}
