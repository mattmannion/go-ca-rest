package controller

import "_/cmd/src/service"

type Deps struct {
	ServiceLayer service.ServiceLayer
}

type ControllerLayer struct {
	PostController IPostController
}

// This is a collection of all the controllers
func NewControllerLayer(dep Deps) *ControllerLayer {
	return &ControllerLayer{
		PostController: NewPostController(dep.ServiceLayer.PostService),
	}
}
