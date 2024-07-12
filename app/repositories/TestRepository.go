package repositories

import (
	"github.com/uzzalhcse/go-fiber-kit/app/models"
	"gorm.io/gorm"
)

type TestRepository struct {
	DB *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{db}
}

// GetAll returns all records from the model
func (r *TestRepository) GetAll() ([]models.TestModel, error) {
	var models []models.TestModel
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
