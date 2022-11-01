package pg_repo

import (
	"_/src/models"
	"_/src/types/repo_types"

	"gorm.io/gorm"
)

type PostRepo struct {
	Db *gorm.DB
}

func NewPostRepo(db *gorm.DB) repo_types.IPostRepo {
	return &PostRepo{Db: db}
}

func (pr *PostRepo) Insert(post *models.Post) (*models.Post, error) {
	pr.Db.Create(&post)

	return post, nil
}

func (pr *PostRepo) GetAll() ([]models.Post, error) {
	posts := []models.Post{}

	pr.Db.Find(&posts)

	return posts, nil
}
