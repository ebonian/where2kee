package controller

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetToilet(c *fiber.Ctx) error {
	toiletRepo := repository.NewToiletRepository()

	toilets, err := toiletRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(toilets)
}

func GetToiletByBuildingID(c *fiber.Ctx) error {
	toiletRepo := repository.NewToiletRepository()

	buildingID := c.Params("buildingID")

	toilets, err := toiletRepo.FindByBuildingID(buildingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(toilets)
}

func CreateToilet(c *fiber.Ctx) error {
	toiletRepo := repository.NewToiletRepository()

	var toiletInput model.CreateToiletInput

	if err := c.BodyParser(&toiletInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	toilet := model.Toilet{
		Name:       toiletInput.Name,
		Image:      toiletInput.Image,
		BuildingID: uuid.MustParse(toiletInput.BuildingID),
	}

	createdToilet, err := toiletRepo.Create(toilet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(createdToilet)
}
