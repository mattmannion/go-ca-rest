package main

import (
	"_/cmd/src/controller"
	"_/cmd/src/repo"
	"_/cmd/src/router"
	"_/cmd/src/service"
	"encoding/json"
	"net/http"
)

var pr repo.PostRepo = repo.NewFirestoreRepo()
var ps service.PostService = service.NewPostService(pr)
var pc controller.PostController = controller.NewPostController(ps)
var r router.Router = router.NewMuxRouter()

func main() {
	r.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-type", "application-json")

		resp.WriteHeader(http.StatusOK)
		json.NewEncoder(resp).Encode(map[string]string{"message": "hello"})
	})
	r.Get("/", pc.GetPosts)
	r.Post("/", pc.PostPost)

	r.Serve(":7890")
}
