package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app/middleware"
	"github.com/jadahbakar/asastarealty-backend/app/response"
	"github.com/jadahbakar/asastarealty-backend/internal/health"
	"github.com/jadahbakar/asastarealty-backend/internal/master/bod"
	"github.com/jadahbakar/asastarealty-backend/pkg/config"
)

type App struct {
	engine *fiber.App
	// config *config.Config
	// database *sql.DB
	logger *os.File
}

func FiberConfig(config *config.Config) *fiber.Config {
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

func FiberLogger(config *config.Config) *os.File {
	// Fiber Create File Logger.
	file, err := os.OpenFile(fmt.Sprintf("%s%s", config.AppLogFolder, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

func createMonolith(engine *fiber.App, db *sql.DB, config *config.Config) {
	router := engine.Group(fmt.Sprintf("%s%s", config.AppURLGroup, config.AppURLVersion))
	// Health Checking
	health.AddRoutes(router)
	// Master BOD
	bodRepo := bod.NewBodRepository(db)
	bodService := bod.NewBodService(bodRepo)
	bod.NewBodHandler(router, bodService)
}

func New(config *config.Config, db *sql.DB) *App {
	// Fiber setup.
	fiberConfig := FiberConfig(config)
	fiberLogger := FiberLogger(config)
	engine := fiber.New(*fiberConfig)
	middleware.FiberMiddleware(engine, fiberLogger)
	createMonolith(engine, db, config)

	// return &App{engine, config, fiberLogger}
	return &App{engine, fiberLogger}

}

// GetEngine
func (app *App) GetEngine() *fiber.App {
	return app.engine
}

// Engine Logger
func (app *App) GetLogger() *os.File {
	return app.logger
}

// // Get Config
// func (app *App) GetConfig() *config.Config {
// 	return app.config
// }
