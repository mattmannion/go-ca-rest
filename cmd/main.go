package main

import (
	"_/cmd/src/controller"
	"_/cmd/src/repo"
	"_/cmd/src/router"
	"_/cmd/src/service"
)

var pr repo.PostRepo = repo.NewFirestoreRepo()
var ps service.PostService = service.NewPostService(pr)
var pc controller.PostController = controller.NewPostController(ps)
var r router.Router = router.NewMuxRouter()

func main() {
	r.Get("/", pc.GetPosts)
	r.Post("/", pc.PostPost)

	r.Serve(":7890")
}
