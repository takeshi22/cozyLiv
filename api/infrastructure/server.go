package infrastructure

import (
	"fmt"
	"log"

	"github.com/takeshi22/cozyLiv/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDatabase() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s TimeZone=%s  port=5432",
		"db",
		"cozyLiv",
		"cozyLiv",
		"cozyLiv",
		"Asia/Tokyo")

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}

	Db = database

	Db.AutoMigrate(&entity.Prefecture{})
	Db.AutoMigrate(&entity.City{})

	return nil
}
