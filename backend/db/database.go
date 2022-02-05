package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-our-schedule/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Loading env file Error: ", err)
		return err
	}

	DB, err = gorm.Open(mysql.Open(os.Getenv("DBDNS")), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connecting Error: ", err)
		return err
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Group{},
		&models.Member{},
		&models.Schedule{},
	)
	if err != nil {
		fmt.Println("DB Connecting Error: ", err)
		return err
	}

	return nil
}

