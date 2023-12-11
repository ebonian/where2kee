package route

import (
	"github.com/gofiber/fiber/v2"
)

func GeneralRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Where2Kee API is running",
		})
	})

}

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Not Found",
			})
		},
	)
}
