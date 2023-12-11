package route

import (
	"github.com/ebonian/where2kee/server/app/controller"
	"github.com/gofiber/fiber/v2"
)

func ToiletRoute(app *fiber.App) {
	toilet := app.Group("/toilet")
	toilet.Get("/", controller.GetToilet)
	toilet.Get("/:id", controller.GetToiletByID)
}
