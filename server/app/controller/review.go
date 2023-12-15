package controller

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateReview(c *fiber.Ctx) error {
	reviewRepo := repository.NewReviewRepository()

	var reviewInput model.CreateReviewInput
	if err := c.BodyParser(&reviewInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	review := model.Review{
		ID:       uuid.New(),
		Username: reviewInput.Username,
		Score:    reviewInput.Score,
		Comment:  reviewInput.Comment,
		ToiletID: uuid.MustParse(reviewInput.ToiletID),
	}

	createdReview, err := reviewRepo.Create(review)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(createdReview)
}
