package models

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name  string
	Email string
	// Add your fields here
}
