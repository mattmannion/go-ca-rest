package service

import "_/cmd/src/repo/types"

type Deps struct {
	PostRepo types.IPostRepo
}

type ServiceLayer struct {
	PostService IPostService
}

// This is a collection of all the services
func NewServiceLayer(dep Deps) *ServiceLayer {
	return &ServiceLayer{
		PostService: NewPostService(dep.PostRepo),
	}
}
