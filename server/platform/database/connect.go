package database

import (
	"github.com/ebonian/where2kee/server/app/model"
	"github.com/ebonian/where2kee/server/pkg/config"
	"github.com/ebonian/where2kee/server/platform/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var log = logger.GetLogger()

func ConnectDB() {
	var err error

	dbConfig := config.DBCfg()

	DB, err = gorm.Open(postgres.Open(dbConfig.ConnStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can't connect to database. error: %v", err)
	}
}

func MigrateDB() {

	err := DB.AutoMigrate(
		&model.Building{},
		&model.Toilet{},
		&model.Review{},
	)

	if err != nil {
		log.Fatalf("Can't migrate database. error: %v", err)
	}

	log.Infoln("Database migrated")
}

func InitializeDB() {
	ConnectDB()
	MigrateDB()

	log.Infoln("Database initialized")
}
