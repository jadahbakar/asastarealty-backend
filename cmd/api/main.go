package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	// new

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app"
	"github.com/jadahbakar/asastarealty-backend/internal/health"
	"github.com/jadahbakar/asastarealty-backend/internal/master/bod"
	"github.com/jadahbakar/asastarealty-backend/pkg/config"
	"github.com/jadahbakar/asastarealty-backend/pkg/utils"
	"github.com/jadahbakar/asastarealty-backend/pkg/version"
)

// https://gist.github.com/rnyrnyrny/282fe705d6e8dc012e482582d7c8ec0b

func init() {
	ver := flag.Bool("version", false, "print version information")
	v := flag.Bool("v", false, "print version information")

	flag.Parse()
	if *v || *ver {
		fmt.Println(version.String())
		os.Exit(0)
	}
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

func main() {
	// // Define Config
	// config, err := config.New()
	// if err != nil {
	// 	log.Printf("error Loading Config -> %v\n", err)
	// }

	// Define Server
	// apps := app.New(config)
	apps := app.SetupApp()
	engine := apps.GetEngine()
	logger := apps.GetLogger()
	config := apps.GetConfig()
	db := apps.GetDB()
	createMonolith(engine, db, apps.GetConfig())

	// Start the Server
	utils.StartFiberWithGracefulShutdown(engine, db, config, logger)
}
