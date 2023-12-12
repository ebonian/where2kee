package data

import (
	"math/rand"

	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/app/repository"
	"github.com/google/uuid"
)

type MockReview struct {
	Username string
	Comment  string
	Score    uint
}

var reviews = []MockReview{
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นห unpleasant",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่สภาพดี แต่บางทีเสียดายที่ไม่มีกระจก",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่แรงมาก ต้องการการดูแลเพิ่มเติม",
		Score:    2,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่เพียงพอ ควรมีการบำรุงรักษาเพิ่มเติม",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นหอมชุ่ย น่าเสียดายที่ไม่มีบูทสเปรย์",
		Score:    4,
	},
	{
		Username: "6538068821",
		Comment:  "บริการทำความสะอาดดีมาก สะอาดสะอ้าน",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่มีกลิ่นหอม ใช้ได้",
		Score:    4,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษาอย่างใกล้ชิด",
		Score:    3,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นห unpleasant",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่สภาพดี แต่บางทีเสียดายที่ไม่มีกระจก",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่แรงมาก ต้องการการดูแลเพิ่มเติม",
		Score:    2,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่เพียงพอ ควรมีการบำรุงรักษาเพิ่มเติม",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นหอมชุ่ย น่าเสียดายที่ไม่มีบูทสเปรย์",
		Score:    4,
	},
	{
		Username: "6538068821",
		Comment:  "บริการทำความสะอาดดีมาก สะอาดสะอ้าน",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่มีกลิ่นหอม ใช้ได้",
		Score:    4,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษาอย่างใกล้ชิด",
		Score:    3,
	},
	{
		Username: "6538078921",
		Comment:  "ห้องน้ำสะอาดและน่าใช้งาน",
		Score:    4,
	},
	{
		Username: "6538240121",
		Comment:  "ทิชชู่มีสภาพดี แต่ควรมีการเพิ่มความสะอาด",
		Score:    3,
	},
	{
		Username: "6538252321",
		Comment:  "กลิ่นฉี่เล็กน้อย ต้องระวังการใช้งาน",
		Score:    2,
	},
	{
		Username: "6538084521",
		Comment:  "บริการทำความสะอาดดีมาก แต่ทิชชู่ใช้ไม่สะดวก",
		Score:    4,
	},
	{
		Username: "6538263721",
		Comment:  "ห้องน้ำไม่สะอาด แนะนำให้ทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538090921",
		Comment:  "ทิชชู่มีกลิ่นหอม พนักงานน่ารัก",
		Score:    5,
	},
	{
		Username: "6538275121",
		Comment:  "กลิ่นฉี่ไม่มี ทิชชู่ใช้ได้ดี",
		Score:    4,
	},
	{
		Username: "6538107321",
		Comment:  "บรรยากาศดี ห้องน้ำสะอาดมาก",
		Score:    5,
	},
	{
		Username: "6538286521",
		Comment:  "กลิ่นหอมชุ่ยมีบ้าง แต่ทำความสะอาดดี",
		Score:    3,
	},
	{
		Username: "6538118921",
		Comment:  "ทิชชู่สภาพดี แต่ควรมีการบำรุงรักษา",
		Score:    2,
	},
	{
		Username: "6538124521",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นหอมชุ่ย",
		Score:    5,
	},
	{
		Username: "6538292121",
		Comment:  "ทิชชู่สภาพดี แต่ต้องระวังความสะอาด",
		Score:    3,
	},
	{
		Username: "6538303521",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษา",
		Score:    2,
	},
	{
		Username: "6538136121",
		Comment:  "บริการทำความสะอาดดีมาก ทิชชู่ใช้งานสะดวก",
		Score:    4,
	},
	{
		Username: "6538317721",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538149321",
		Comment:  "ทิชชู่มีกลิ่นหอม บรรยากาศดี",
		Score:    5,
	},
	{
		Username: "6538320921",
		Comment:  "กลิ่นฉี่ไม่มี ทิชชู่ใช้ได้ดี",
		Score:    4,
	},
	{
		Username: "6538154521",
		Comment:  "บรรยากาศดี ห้องน้ำสะอาดมาก",
		Score:    5,
	},
	{
		Username: "6538336121",
		Comment:  "กลิ่นหอมชุ่ยมีบ้าง แต่ทำความสะอาดดี",
		Score:    3,
	},
	{
		Username: "6538166921",
		Comment:  "ทิชชู่สภาพดี แต่ควรมีการบำรุงรักษา",
		Score:    2,
	},
	{
		Username: "6538178121",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นหอมชุ่ย",
		Score:    5,
	},
	{
		Username: "6538349721",
		Comment:  "ทิชชู่สภาพดี แต่ต้องระวังความสะอาด",
		Score:    3,
	},
	{
		Username: "6538351321",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษา",
		Score:    2,
	},
	{
		Username: "6538182921",
		Comment:  "บริการทำความสะอาดดีมาก ทิชชู่ใช้งานสะดวก",
		Score:    4,
	},
	{
		Username: "6538364521",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538196121",
		Comment:  "ทิชชู่มีกลิ่นหอม บรรยากาศดี",
		Score:    5,
	},
	{
		Username: "6538377721",
		Comment:  "กลิ่นฉี่ไม่มี ทิชชู่ใช้ได้ดี",
		Score:    4,
	},
	{
		Username: "6538209321",
		Comment:  "บรรยากาศดี ห้องน้ำสะอาดมาก",
		Score:    5,
	},
	{
		Username: "6538380921",
		Comment:  "กลิ่นหอมชุ่ยมีบ้าง แต่ทำความสะอาดดี",
		Score:    3,
	},
	{
		Username: "6538214521",
		Comment:  "ทิชชู่สภาพดี แต่ควรมีการบำรุงรักษา",
		Score:    2,
	},
	{
		Username: "6538178121",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นหอมชุ่ย",
		Score:    5,
	},
	{
		Username: "6538349721",
		Comment:  "ทิชชู่สภาพดี แต่ต้องระวังความสะอาด",
		Score:    3,
	},
	{
		Username: "6538351321",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษา",
		Score:    2,
	},
	{
		Username: "6538182921",
		Comment:  "บริการทำความสะอาดดีมาก ทิชชู่ใช้งานสะดวก",
		Score:    4,
	},
	{
		Username: "6538364521",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538196121",
		Comment:  "ทิชชู่มีกลิ่นหอม บรรยากาศดี",
		Score:    5,
	},
	{
		Username: "6538377721",
		Comment:  "กลิ่นฉี่ไม่มี ทิชชู่ใช้ได้ดี",
		Score:    4,
	},
	{
		Username: "6538209321",
		Comment:  "บรรยากาศดี ห้องน้ำสะอาดมาก",
		Score:    5,
	},
	{
		Username: "6538380921",
		Comment:  "กลิ่นหอมชุ่ยมีบ้าง แต่ทำความสะอาดดี",
		Score:    3,
	},
	{
		Username: "6538214521",
		Comment:  "ทิชชู่สภาพดี แต่ควรมีการบำรุงรักษา",
		Score:    2,
	},
	{
		Username: "6538228121",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นหอมชุ่ย",
		Score:    5,
	},
	{
		Username: "6538399721",
		Comment:  "ทิชชู่สภาพดี แต่ต้องระวังความสะอาด",
		Score:    3,
	},
	{
		Username: "6538401321",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษา",
		Score:    2,
	},
	{
		Username: "6538232921",
		Comment:  "บริการทำความสะอาดดีมาก ทิชชู่ใช้งานสะดวก",
		Score:    4,
	},
	{
		Username: "6538414521",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538246121",
		Comment:  "ทิชชู่มีกลิ่นหอม บรรยากาศดี",
		Score:    5,
	},
	{
		Username: "6538427721",
		Comment:  "กลิ่นฉี่ไม่มี ทิชชู่ใช้ได้ดี",
		Score:    4,
	},
	{
		Username: "6538259321",
		Comment:  "บรรยากาศดี ห้องน้ำสะอาดมาก",
		Score:    5,
	},
	{
		Username: "6538430921",
		Comment:  "กลิ่นหอมชุ่ยมีบ้าง แต่ทำความสะอาดดี",
		Score:    3,
	},
	{
		Username: "6538262121",
		Comment:  "ทิชชู่สภาพดี แต่ควรมีการบำรุงรักษา",
		Score:    2,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นห unpleasant",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่สภาพดี แต่บางทีเสียดายที่ไม่มีกระจก",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่แรงมาก ต้องการการดูแลเพิ่มเติม",
		Score:    2,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่เพียงพอ ควรมีการบำรุงรักษาเพิ่มเติม",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นหอมชุ่ย น่าเสียดายที่ไม่มีบูทสเปรย์",
		Score:    4,
	},
	{
		Username: "6538068821",
		Comment:  "บริการทำความสะอาดดีมาก สะอาดสะอ้าน",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่มีกลิ่นหอม ใช้ได้",
		Score:    4,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษาอย่างใกล้ชิด",
		Score:    3,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำสะอาดมาก ไม่มีกลิ่นห unpleasant",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่สภาพดี แต่บางทีเสียดายที่ไม่มีกระจก",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่แรงมาก ต้องการการดูแลเพิ่มเติม",
		Score:    2,
	},
	{
		Username: "6538068821",
		Comment:  "ห้องน้ำไม่สะอาด ควรทำความสะอาดบ่อยขึ้น",
		Score:    1,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่เพียงพอ ควรมีการบำรุงรักษาเพิ่มเติม",
		Score:    3,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นหอมชุ่ย น่าเสียดายที่ไม่มีบูทสเปรย์",
		Score:    4,
	},
	{
		Username: "6538068821",
		Comment:  "บริการทำความสะอาดดีมาก สะอาดสะอ้าน",
		Score:    5,
	},
	{
		Username: "6538231721",
		Comment:  "ทิชชู่มีกลิ่นหอม ใช้ได้",
		Score:    4,
	},
	{
		Username: "6538223721",
		Comment:  "กลิ่นฉี่เล็กน้อย ควรดูแลรักษาอย่างใกล้ชิด",
		Score:    3,
	},
	{
		Username: "6538078921",
		Comment:  "ห้องน้ำสะอาดและน่าใช้งาน",
		Score:    4,
	},
	{
		Username: "6538240121",
		Comment:  "ทิชชู่มีสภาพดี แต่ควรมีการเพิ่มความสะอาด",
		Score:    3,
	},
	{
		Username: "6538252321",
		Comment:  "กลิ่นฉี่เล็กน้อย ต้องระวังการใช้งาน",
		Score:    2,
	},
}

func CreateReview() error {
	reviewRepo := repository.NewReviewRepository()
	toiletRepo := repository.NewToiletRepository()

	toilets, err := toiletRepo.FindAll()
	if err != nil {
		return err
	}

	for _, review := range reviews {
		_, err := reviewRepo.Create(model.Review{
			ID:       uuid.New(),
			Username: review.Username,
			Comment:  review.Comment,
			Score:    review.Score,
			ToiletID: toilets[rand.Intn(len(toilets))].ID,
		})
		if err != nil {
			return err
		}
	}

	return nil
}
