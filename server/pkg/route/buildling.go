package route

import (
	"github.com/ebonian/where2kee/server/app/controller"
	"github.com/gofiber/fiber/v2"
)

func BuildingRoute(app *fiber.App) {
	building := app.Group("/building")

	building.Get("/", controller.GetBuilding)
	building.Get("/:id", controller.GetBuildingByID)
}
