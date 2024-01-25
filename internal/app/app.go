package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ARVG9866/new_project/dev"
	"github.com/ARVG9866/new_project/internal/models"
	"github.com/ARVG9866/new_project/internal/service"
	"github.com/ARVG9866/new_project/internal/storage"

	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"

	_ "github.com/lib/pq"
)

type App struct {
	appConfig *models.Config

	server *Server
	db     *DataBase
	//service *Service
}

type Server struct {
	handler *http.Handler
	logger  *slog.Logger

	service service.IService
}

type DataBase struct {
	db     *sqlx.DB
	logger *slog.Logger
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	a.setConfig()
	a.initDB()

	a.server = a.initServer()

	return a, nil
}

func (a *App) setConfig() error {
	// Use local config
	err := dev.SetConfig()
	if err != nil {
		return err
	}

	conf := models.Config{}

	envconfig.MustProcess("", &conf)
	a.appConfig = &conf

	return nil
}

func (a *App) initServer() *Server {
	storage := storage.NewStorage(a.db.db)
	service := service.NewService(storage)

	return &Server{
		logger:  slog.Default().With("name", "server"),
		service: service,
	}
}

func (a *App) initDB() {
	a.db = &DataBase{
		logger: slog.Default().With("name", "db"),
	}

	sqlConnectionString := a.getSqlConnectionString()

	var err error
	a.db.db, err = sqlx.Open("postgres", sqlConnectionString)

	if err != nil {
		a.db.logger.Error("failed to opening connection to db: ", err.Error())
	}

	if err = a.db.db.Ping(); err != nil {
		a.db.logger.Error("failed to connect to the database: ", err.Error())
	}

}

func (a *App) getSqlConnectionString() string {
	sqlConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%v",
		a.appConfig.DB.User,
		a.appConfig.DB.Password,
		a.appConfig.DB.Host,
		a.appConfig.DB.Port,
		a.appConfig.DB.Database,
		a.appConfig.DB.SSLMode,
	)

	return sqlConnectionString
}
