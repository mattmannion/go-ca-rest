package service_types

import (
	"_/src/models"
	"_/src/types/repo_types"
)

type IPostService interface {
	Validate(post *models.Post) error
	FindAll() ([]models.Post, error)
	Create(post *models.Post) (*models.Post, error)
}

type INewPostService interface {
	NewPostService(PostRepo repo_types.IPostRepo) IPostService
}
