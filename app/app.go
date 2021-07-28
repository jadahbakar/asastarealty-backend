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
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/bod"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/health"
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
		AppName:       config.App.Name,
		Prefork:       config.App.Prefork,
		CaseSensitive: config.App.CaseSensitive,
		ReadTimeout:   time.Second * time.Duration(config.App.ReadTimeOut),
		WriteTimeout:  time.Second * time.Duration(config.App.WriteTimeOut),
		ErrorHandler:  response.DefaultErrorHandler,
	}
}

// func FiberLogger(config *config.Config) *os.File {
func FiberLogger(loggerPath string) *os.File {
	// Fiber Create File Logger.
	// file, err := os.OpenFile(fmt.Sprintf("%s%s", config.App.LogFolder, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	file, err := os.OpenFile(fmt.Sprintf("%s%s", loggerPath, "fiber.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	return file
}

func createMonolith(engine *fiber.App, dbClient *sql.DB, config *config.Config) {
	router := engine.Group(fmt.Sprintf("%s%s", config.App.URLGroup, config.App.URLVersion))
	// Health Checking
	health.AddRoutes(router)
	// Master BOD
	bodRepo := bod.NewBodRepository(dbClient)
	bodService := bod.NewBodService(bodRepo)
	bod.NewBodHandler(router, bodService)

}

func New(config *config.Config, dbClient *sql.DB) *App {
	// Fiber setup.
	fiberConfig := FiberConfig(config)
	// fiberLogger := FiberLogger(config)
	fiberLogger := FiberLogger(config.App.LogFolder)

	engine := fiber.New(*fiberConfig)
	middleware.FiberMiddleware(engine, fiberLogger)
	createMonolith(engine, dbClient, config)

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
