package data

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/google/uuid"
)

var BuildingId = map[string]uuid.UUID{
	"ตึก 1":   uuid.MustParse("a8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
	"ตึก 2":   uuid.MustParse("b8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
	"ตึก 3":   uuid.MustParse("c8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
	"ตึก 4":   uuid.MustParse("d8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
	"ตึก 100": uuid.MustParse("e8a3b8c0-6b1a-4b7e-8f1a-9e6b9b9b9b9b"),
}

var buildings = []model.Building{
	{
		ID:      BuildingId["ตึก 1"],
		Name:    "ตึก 1",
		Toilets: []model.Toilet{},
	},
	{
		ID:      BuildingId["ตึก 2"],
		Name:    "ตึก 2",
		Toilets: []model.Toilet{},
	},
	{
		ID:      BuildingId["ตึก 3"],
		Name:    "ตึก 3",
		Toilets: []model.Toilet{},
	},
	{
		ID:      BuildingId["ตึก 4"],
		Name:    "ตึก 4",
		Toilets: []model.Toilet{},
	},
	{
		ID:      BuildingId["ตึก 100"],
		Name:    "ตึก 100",
		Toilets: []model.Toilet{},
	},
}

func CreateBuilding() error {
	buildingRepo := repository.NewBuildingRepository()

	for _, building := range buildings {
		_, err := buildingRepo.Create(building)
		if err != nil {
			return err
		}
	}

	return nil
}
