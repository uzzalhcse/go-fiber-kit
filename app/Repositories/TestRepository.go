package Repositories

import (
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"gorm.io/gorm"
)

type TestRepository struct {
	DB *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{db}
}

// GetAll returns all records from the model
func (r *TestRepository) GetAll() ([]Models.TestModel, error) {
	var models []Models.TestModel
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
