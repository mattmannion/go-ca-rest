package post_controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"_/cmd/src/controller/controller_types"
	"_/cmd/src/models"
	"_/cmd/src/service/service_types"
)

type controller struct{ PostService service_types.IPostService }

func NewPostController(PostService service_types.IPostService) controller_types.IPostController {
	return &controller{PostService: PostService}
}

func (c *controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := c.PostService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (c *controller) PostPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	post := &models.Post{}

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	err = c.PostService.Validate(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{"error": fmt.Sprint(err)})
		return
	}

	res, err := c.PostService.Create(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{"error": fmt.Sprint(err)})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(res)
}
