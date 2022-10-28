package service_types

import (
	"_/cmd/src/models"
	"_/cmd/src/repo/repo_types"
)

type IPostService interface {
	Validate(post *models.Post) error
	Create(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type INewPostService interface {
	NewPostService(PostRepo repo_types.IPostRepo) IPostService
}
