package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/asastarealty-backend/app/response"
)

func AddRoutes(router fiber.Router) {
	router.Get("/health", GetHealth)
}

func GetHealth(c *fiber.Ctx) error {
	// Return status 200 OK.
	return response.NewSuccess(c, fiber.StatusOK, "healthty", nil)
}
