package data

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/google/uuid"
)

var toilets = []model.Toilet{
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 3",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 4",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 4",
		Image:      "",
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น M",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น M",
		Image:      "",
		BuildingID: BuildingId["ตึก 4"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น M",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น M",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2 ฝั่ง 201A",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2 ฝั่ง 201A",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2 ฝั่ง 201B",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2 ฝั่ง 201B",
		Image:      "",
		BuildingID: BuildingId["ตึก 100"],
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
