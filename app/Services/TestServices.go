package Services

import (
	"github.com/uzzalhcse/amadeus-go/app/Models"
	"github.com/uzzalhcse/amadeus-go/app/Repositories"
)

type TestService struct {
	Repository *Repositories.TestRepository
}

func NewTestService(repo *Repositories.TestRepository) *TestService {
	return &TestService{Repository: repo}
}

// GetAll returns all records from the model using the repository
func (s *TestService) GetAll() ([]Models.TestModel, error) {
	return s.Repository.GetAll()
}
