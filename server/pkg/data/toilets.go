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
		Image:      ToiletImages["men/1/1"],
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      ToiletImages["women/1/1"],
		BuildingID: BuildingId["ตึก 1"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      ToiletImages["men/2/1"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      ToiletImages["women/2/1"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      ToiletImages["men/2/2"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      ToiletImages["women/2/2"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 3",
		Image:      ToiletImages["men/2/3"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 3",
		Image:      ToiletImages["women/2/3"],
		BuildingID: BuildingId["ตึก 2"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      ToiletImages["men/3"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1",
		Image:      ToiletImages["women/3"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1 ฝั่ง TrueLab",
		Image:      ToiletImages["men/3/true"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 1 ฝั่ง TrueLab",
		Image:      ToiletImages["women/3/true"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      ToiletImages["men/3"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
		Image:      ToiletImages["women/3"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 3",
		Image:      ToiletImages["women/3/old"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 4",
		Image:      ToiletImages["men/3/old"],
		BuildingID: BuildingId["ตึก 3"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 1",
		Image:      ToiletImages["men/4/1"],
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
		Name:       "ห้องน้ำชายชั้น M",
		Image:      ToiletImages["men/4/m"],
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
		Image:      ToiletImages["men/100/m"],
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น M",
		Image:      ToiletImages["women/4/m"],
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำชายชั้น 2",
		Image:      ToiletImages["men/100/2"],
		BuildingID: BuildingId["ตึก 100"],
	},
	{
		ID:         uuid.New(),
		Name:       "ห้องน้ำหญิงชั้น 2",
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
