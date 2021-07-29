package app

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/internal/config"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/bod"
	"github.com/jadahbakar/asastarealty-backend/internal/domain/master/health"
)

// ini tidak bisa di test, karena ini fungsi yang tidak membalikan apapun.
// hanya fungsi untuk meringkas code.
func MonolithIOC(engine *fiber.App, dbClient *sql.DB, config *config.Config) {
	router := engine.Group(fmt.Sprintf("%s%s", config.App.URLGroup, config.App.URLVersion))

	// Health Checking
	health.AddRoutes(router)

	// Master BOD
	bodRepo := bod.NewBodRepository(dbClient)
	bodService := bod.NewBodService(bodRepo)
	bod.NewBodHandler(router, bodService)
}
