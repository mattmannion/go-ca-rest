package repo_types

import "_/src/models"

type IPostRepo interface {
	Insert(post *models.Post) (*models.Post, error)
	GetAll() ([]models.Post, error)
}
