package server

import (
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/ebonian/where2kee/server/pkg/data"
)

func CreateData() {
	buildingRepo := repository.NewBuildingRepository()

	buildings, err := buildingRepo.FindAll()
	if err != nil {
		panic(err)
	}

	if len(buildings) > 0 {
		return
	}

	data.CreateBuilding()
	data.CreateToilet()
	data.CreateReview()
}
