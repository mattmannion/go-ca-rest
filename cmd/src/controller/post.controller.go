package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"_/cmd/src/model"
	"_/cmd/src/service"
)

type IPostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	PostPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct {
	PostService service.IPostService
}

func NewPostController(PostService service.IPostService) IPostController {
	return &controller{
		PostService: PostService,
	}
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

	post := &model.Post{}

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
