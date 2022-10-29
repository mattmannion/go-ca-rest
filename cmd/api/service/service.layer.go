package service

import (
	"_/cmd/api/repo"
	"_/cmd/api/service/post_service"
	"_/cmd/types/service_types"
)

type Deps struct {
	RepoLayer repo.RepoLayer
}

type ServiceLayer struct{ PostService service_types.IPostService }

func NewServiceLayer(deps Deps) *ServiceLayer {
	return &ServiceLayer{
		PostService: post_service.NewPostService(deps.RepoLayer.PostRepo),
	}
}
