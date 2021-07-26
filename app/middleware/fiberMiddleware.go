package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"
)

// https://github.com/gofiber/fiber/issues/1293
// https://github.com/gofiber/fiber/issues/683
// https://www.programmersought.com/article/24337694346/

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App, logFile *os.File) {
	//----- Logger
	ConfigLogger := logger.Config{
		Next:   nil,
		Format: `{"pid":${pid}, "timestamp":"${time}", "status":${status}, "latency":"${latency}", "method":"${method}", "path":"${path}"}` + "\n",
		// Format:     "[${time}] ${pid} | ${locals:requestid} | ${ip}| ${status} - ${latency} | ${method} | ${path}\n",
		TimeFormat: "2006/Jan/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
		Output:     logFile,
	}
	//----- CORS
	ConfigCORS := cors.Config{
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	//----- Compress
	ConfigCompress := compress.Config{
		Next:  nil,
		Level: compress.LevelBestSpeed,
	}
	//----- CSRF
	ConfigCSRF := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token",
		CookieName:     "csrf_",
		CookieSameSite: "Strict",
		Expiration:     1 * time.Hour,
		KeyGenerator:   utils.UUID,
		CookieSecure:   true,
		CookieHTTPOnly: true,
	}

	a.Use(
		requestid.New(),
		// Add CORS to each route.
		cors.New(ConfigCORS),
		// Add logger.
		logger.New(ConfigLogger),
		// Add CSRF
		csrf.New(ConfigCSRF),
		// Add Compress
		compress.New(ConfigCompress),
	)
}
