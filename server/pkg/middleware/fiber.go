package middleware

import (
	"github.com/ebonian/where2kee/server/platform/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var log = logger.GetLogger()

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
			AllowMethods: "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		}),

		RequestLogger(log),
	)
}
