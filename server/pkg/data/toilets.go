package data

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/google/uuid"
)

var toilets = []model.Toilet{
	{
		ID:         uuid.MustParse("a8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.MustParse("b8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
}

func CreateToilet() error {
	toiletRepo := repository.NewToiletRepository()

	for _, toilet := range toilets {
		_, err := toiletRepo.Create(toilet)
		if err != nil {
			return err
		}
	}

	return nil
}
