package repo_types

import "_/cmd/src/models"

type IPostRepo interface {
	Save(post *models.Post) (*models.Post, error)
	FindAll() ([]models.Post, error)
}

type INewPostRepo interface {
	// interface{} is used here to allow us to
	// create new repos with varying dependencies
	// aslong as the constructor returns IPostRepo
	NewPostRepo(interface{}) IPostRepo
}
