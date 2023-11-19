package Models

import "gorm.io/gorm"

type TestModel struct {
	gorm.Model
	Name string
	// Add your fields here
}
