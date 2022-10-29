package main

import (
	"_/cmd/api/controller"
	"_/cmd/api/repo"
	"_/cmd/api/router"
	"_/cmd/api/service"
	"_/cmd/constants"
	"encoding/json"
	"net/http"
)

var (
	RepoLayer  = repo.NewRepoLayer()
	Services   = service.NewServiceLayer(service.Deps{RepoLayer: *RepoLayer})
	Contollers = controller.NewControllerLayer(controller.Deps{ServiceLayer: *Services})
	Router     = router.NewMuxRouter()
)

func main() {
	Router.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application-json")

		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(map[string]string{"message": "hello"})
	})

	Router.Get(constants.ApiPrefixV1+"/posts", Contollers.PostController.GetPosts)
	Router.Post(constants.ApiPrefixV1+"/posts", Contollers.PostController.PostPost)

	Router.Serve(":7890")
}
