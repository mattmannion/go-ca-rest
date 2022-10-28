package service

import (
	"_/cmd/src/model"
	"_/cmd/src/repo/types"
	"errors"
	"math/rand"
)

type IPostService interface {
	Validate(post *model.Post) error
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

type PostService struct {
	PostRepo types.IPostRepo
}

func NewPostService(PostRepo types.IPostRepo) IPostService {
	return &PostService{
		PostRepo: PostRepo,
	}
}

func (*PostService) Validate(post *model.Post) error {
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

func (s *PostService) Create(post *model.Post) (*model.Post, error) {
	post.Id = rand.Int()
	return s.PostRepo.Save(post)
}

func (s *PostService) FindAll() ([]model.Post, error) {
	return s.PostRepo.FindAll()
}
