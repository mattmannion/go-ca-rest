package mock_services

import (
	"_/cmd/models"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockPostService struct {
	mock.Mock
	ReturnPost func(*models.Post) *models.Post
}

func (m *MockPostService) Validate(post *models.Post) error {
	args := m.Called()
	args.Get(1)
	return args.Error(1)
}

func (m *MockPostService) Create(post *models.Post) (*models.Post, error) {
	args := m.Called(post)
	fmt.Printf("here %v\n", post)
	return m.ReturnPost(post), args.Error(1)
}

func (m *MockPostService) FindAll() ([]models.Post, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Post), args.Error(1)
}
