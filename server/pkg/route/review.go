package route

import (
	"github.com/ebonian/where2kee/server/app/controller"
	"github.com/gofiber/fiber/v2"
)

func ReviewRoute(app *fiber.App) {
	review := app.Group("/review")

	review.Post("/", controller.CreateReview)
}
