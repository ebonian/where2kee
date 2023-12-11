package controller

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetBuilding(c *fiber.Ctx) error {
	buildingRepo := repository.NewBuildingRepository()

	buildings, err := buildingRepo.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(buildings)
}

func GetBuildingByID(c *fiber.Ctx) error {
	buildingRepo := repository.NewBuildingRepository()

	building, err := buildingRepo.FindByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(building)
}

func CreateBuilding(c *fiber.Ctx) error {
	buildingRepo := repository.NewBuildingRepository()

	var buildingInput model.CreateBuildingInput
	if err := c.BodyParser(&buildingInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	building := model.Building{
		ID:      uuid.New(),
		Name:    buildingInput.Name,
		Toilets: []model.Toilet{},
	}

	createdBuilding, err := buildingRepo.Create(building)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(createdBuilding)
}
