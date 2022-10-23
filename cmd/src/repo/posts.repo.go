package repo

import "_/cmd/src/model"

type PostRepo interface {
	Save(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}
