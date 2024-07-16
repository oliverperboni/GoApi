// config/config.go
package config

import (
	"github.com/oliverperboni/GoApi/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// dsn := "user=postgres password=oliver dbname=bookShelf port=5432 sslmode=disable"
	dsn := "host=localhost user=postgres password=oliver dbname=bookShelf port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&schemas.Book{})

	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
