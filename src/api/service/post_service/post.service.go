package post_service

import (
	"_/src/models"
	"_/src/types/repo_types"
	"_/src/types/service_types"
	"errors"
	"math/rand"
)

type PostService struct{ PostRepo repo_types.IPostRepo }

func NewPostService(PostRepo repo_types.IPostRepo) service_types.IPostService {
	return &PostService{PostRepo: PostRepo}
}

func (*PostService) Validate(post *models.Post) error {
	if post == nil {
		return errors.New("no Posts")
	}

	if post.Title == "" {
		return errors.New("no title")
	}

	if post.Text == "" {
		return errors.New("no text")
	}

	return nil
}

func (s *PostService) FindAll() ([]models.Post, error) {
	return s.PostRepo.GetAll()
}

func (s *PostService) Create(post *models.Post) (*models.Post, error) {
	post.Id = rand.Int()
	return s.PostRepo.Insert(post)
}
