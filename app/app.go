package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jadahbakar/asastarealty-backend/app/middleware"
	"github.com/jadahbakar/asastarealty-backend/app/response"
	"github.com/jadahbakar/asastarealty-backend/pkg/config"
)

type App struct {
	engine   *fiber.App
	config   *config.Config
	database *sql.DB
	logger   *os.File
}

func fiberConfig(config *config.Config) *fiber.Config {
	// Fiber config.
	return &fiber.Config{
		AppName:       config.AppName,
		Prefork:       config.AppPrefork,
		CaseSensitive: config.AppCaseSensitive,
		ReadTimeout:   time.Second * time.Duration(config.AppReadTimeOut),
		WriteTimeout:  time.Second * time.Duration(config.AppWriteTimeOut),
		ErrorHandler:  response.DefaultErrorHandler,
	}
}

func fiberLogger(config *config.Config) *os.File {
	// Fiber Create File Logger.
	file, err := os.OpenFile(fmt.Sprintf("%s%s", config.AppLogFolder, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

func SetupApp() *App {
	// Define Config
	config, err := config.New()
	if err != nil {
		log.Printf("error Loading Config -> %v\n", err)
	} // Define Server
	apps := New(config)
	return apps

}

func New(config *config.Config) *App {
	// Fiber setup.
	fiberConfig := fiberConfig(config)
	fiberLogger := fiberLogger(config)
	engine := fiber.New(*fiberConfig)
	middleware.FiberMiddleware(engine, fiberLogger)

	// Database setup.
	connConfig, _ := pgx.ParseConfig(config.DbUrl)
	connStr := stdlib.RegisterConnConfig(connConfig)
	dbConn, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return &App{engine, config, dbConn, fiberLogger}
}

// GetEngine
func (app *App) GetEngine() *fiber.App {
	return app.engine
}

// Engine Logger
func (app *App) GetLogger() *os.File {
	return app.logger
}

// GetDB
func (app *App) GetDB() *sql.DB {
	return app.database
}

// GetDB
func (app *App) GetConfig() *config.Config {
	return app.config
}
