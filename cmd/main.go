package main

import (
	"_/cmd/src/controller"
	"_/cmd/src/repo"
	"_/cmd/src/router"
	"_/cmd/src/service"
	"encoding/json"
	"net/http"
)

var (
	RepoLayer  = repo.NewRepoLayer()
	Services   = service.NewServiceLayer(service.Deps{PostRepo: RepoLayer.PostRepo})
	Contollers = controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router     = router.NewMuxRouter()
)

func main() {
	Router.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application-json")

		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(map[string]string{"message": "hello"})
	})

	Router.Get("/", Contollers.PostController.GetPosts)
	Router.Post("/", Contollers.PostController.PostPost)

	Router.Serve(":7890")
}
