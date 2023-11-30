package database

import (
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.TestModel{},
		&models.User{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
