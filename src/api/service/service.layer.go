package service

import (
	"_/src/api/repo"
	"_/src/api/service/post_service"
	"_/src/types/service_types"
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
