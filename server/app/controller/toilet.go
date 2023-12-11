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

func GetToiletByID(c *fiber.Ctx) error {
	toiletRepo := repository.NewToiletRepository()

	toilet, err := toiletRepo.FindByID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(toilet)
}

func CreateToilet(c *fiber.Ctx) error {
	toiletRepo := repository.NewToiletRepository()

	var toiletInput model.CreateToiletInput

	if err := c.BodyParser(&toiletInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	toilet := model.Toilet{
		Name:       toiletInput.Name,
		Score:      toiletInput.Score,
		Location:   toiletInput.Location,
		Image:      toiletInput.Image,
		BuildingID: uuid.MustParse(toiletInput.BuildingID),
	}

	createdToilet, err := toiletRepo.Create(toilet)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(createdToilet)
}
