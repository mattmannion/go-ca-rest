package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"_/cmd/src/model"
	"_/cmd/src/service"
)

var post_service service.PostService

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	PostPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController(svc service.PostService) PostController {
	post_service = svc
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := post_service.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) PostPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	post := &model.Post{}

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	err = post_service.Validate(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{"error": fmt.Sprint(err)})
		return
	}

	res, err := post_service.Create(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{"error": fmt.Sprint(err)})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(res)
}
