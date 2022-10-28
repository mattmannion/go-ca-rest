package types

import "_/cmd/src/model"

type IPostRepo interface {
	Save(post *model.Post) (*model.Post, error)
	FindAll() ([]model.Post, error)
}
