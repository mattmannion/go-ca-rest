package post_controller

import (
	"_/src/models"
	"_/src/types/controller_types"
	"_/src/types/service_types"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostController struct{ PostService service_types.IPostService }

func NewPostController(PostService service_types.IPostService) controller_types.IPostController {
	return &PostController{PostService: PostService}
}

func (c *PostController) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := c.PostService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{"error": "Could not find Posts"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (c *PostController) PostPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	post := &models.Post{}

	json.NewDecoder(req.Body).Decode(&post)

	err := c.PostService.Validate(post)
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
