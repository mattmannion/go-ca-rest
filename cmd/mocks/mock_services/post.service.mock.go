package mock_services

import (
	"_/cmd/models"

	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
}

func (m *MockPostService) Validate(post *models.Post) error {
	args := m.Called(post)
	return args.Error(0)
}

func (m *MockPostService) Create(post *models.Post) (*models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*models.Post), args.Error(1)
}

func (m *MockPostService) FindAll() ([]models.Post, error) {
	args := m.Called()
	return args.Get(0).([]models.Post), args.Error(1)
}
