package mock_repo

import (
	"_/src/models"

	"github.com/stretchr/testify/mock"
)

type MockPostRepo struct {
	mock.Mock
}

func (m *MockPostRepo) GetAll() ([]models.Post, error) {
	args := m.Called()
	return args.Get(0).([]models.Post), args.Error(1)
}

func (m *MockPostRepo) Insert(post *models.Post) (*models.Post, error) {
	args := m.Called(post)
	return args.Get(0).(*models.Post), args.Error(1)
}
