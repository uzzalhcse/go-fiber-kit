package models

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name string
	// Add your fields here
}
