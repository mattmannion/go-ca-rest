package service

import (
	"_/cmd/src/model"
	"_/cmd/src/repo"
	"errors"
	"math/rand"
)

var posts_repo repo.PostRepo

type PostService interface {
	Validate(post *model.Post) error
	Create(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}

type service struct{}

func NewPostService(pr repo.PostRepo) PostService {
	posts_repo = pr
	return &service{}
}

func (*service) Validate(post *model.Post) error {
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

func (*service) Create(post *model.Post) (*model.Post, error) {
	post.Id = rand.Int()
	return posts_repo.Save(post)

}

func (*service) FindAll() ([]model.Post, error) {
	return posts_repo.FindAll()
}
