package app

import (
	"database/sql"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app/logger"
	"github.com/jadahbakar/asastarealty-backend/app/middleware"
	"github.com/jadahbakar/asastarealty-backend/app/response"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
)

type App struct {
	engine   *fiber.App
	config   *config.Config
	database *sql.DB
	logger   *os.File
}

func New(config *config.Config, dbClient *sql.DB) *App {
	engine := fiber.New(fiber.Config{
		AppName:       config.App.Name,
		Prefork:       config.App.Prefork,
		CaseSensitive: config.App.CaseSensitive,
		ReadTimeout:   time.Second * time.Duration(config.App.ReadTimeOut),
		WriteTimeout:  time.Second * time.Duration(config.App.WriteTimeOut),
		ErrorHandler:  response.DefaultErrorHandler,
	})

	// ini bisa membalikan error, dan saya tidak handle ini karena
	// nanti akan banyak style yang di rubah
	// asumsi saya NewFiberLogger always not error
	fiberLogger, _ := logger.NewFiberLogger(config.App.LogFolder)

	// baris dibawah ini tidak bisa di test, karena ini fungsi yang tidak membalikan apapun, hanya fungsi untuk meringkas code.
	middleware.FiberMiddleware(engine, fiberLogger)

	MonolithIOC(engine, dbClient, config)

	return &App{engine, config, dbClient, fiberLogger}
}

// GetEngine
func (app *App) GetEngine() *fiber.App {
	return app.engine
}

// Engine Logger
func (app *App) GetLogger() *os.File {
	return app.logger
}

// Get Config
func (app *App) GetConfig() *config.Config {
	return app.config
}

// Get Database
func (app *App) GetDB() *sql.DB {
	return app.database
}
