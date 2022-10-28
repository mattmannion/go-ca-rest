package service

import (
	"_/cmd/src/repo"
	"_/cmd/src/service/post_service"
	"_/cmd/src/service/service_types"
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
