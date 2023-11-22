package database

import (
	"fmt"
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&Models.TestModel{},
		&Models.User{},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
