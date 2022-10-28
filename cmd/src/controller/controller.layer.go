package controller

import (
	"_/cmd/src/controller/controller_types"
	"_/cmd/src/controller/post_controller"
	"_/cmd/src/service"
)

type Deps struct{ ServiceLayer service.ServiceLayer }

type ControllerLayer struct {
	PostController controller_types.IPostController
}

func NewControllerLayer(dep Deps) *ControllerLayer {
	return &ControllerLayer{
		PostController: post_controller.NewPostController(dep.ServiceLayer.PostService),
	}
}
