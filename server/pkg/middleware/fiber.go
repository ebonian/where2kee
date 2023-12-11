package middleware

import (
	"github.com/ebonian/where2kee/server/pkg/config"
	"github.com/ebonian/where2kee/server/platform/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var log = logger.GetLogger()

func FiberMiddleware(app *fiber.App) {
	appConfig := config.AppCfg()

	app.Use(
		cors.New(cors.Config{
			AllowOrigins: appConfig.FrontendURL,
			AllowHeaders: "Origin, Content-Type, Accept, Authorization, X-Requested-With",
			AllowMethods: "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		}),

		RequestLogger(log),
	)
}
