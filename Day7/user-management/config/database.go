package config

import (
	"fmt"
	"os"
	"user-management/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/joho/godotenv/autoload"
)

var (
	DB *gorm.DB
	err error
)

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("db_url")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")
}

func Migrate() {
	DB.AutoMigrate(&entity.User{})
}