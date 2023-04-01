package controller

import (
	"_/src/api/controller/post_controller"
	"_/src/api/service"
	"_/src/types/controller_types"
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
