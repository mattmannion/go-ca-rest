package mock_services

import (
	"_/cmd/models"

	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) Validate(post *models.Post) error {
	args := m.Called()
	return args.Error(1)
}

func (m *MockPostService) Create(post *models.Post) (*models.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.(*models.Post), args.Error(1)
}

func (m *MockPostService) FindAll() ([]models.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Post), args.Error(1)
}
