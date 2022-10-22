package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"_/cmd/src/model"
	"_/cmd/src/repo"
)

// var posts []model.Post = []model.Post{{Id: 1, Title: "one", Text: "Hello"}}
var posts_repo repo.PostRepo = repo.NewPostRepo()

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	posts, err := posts_repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func PostPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")

	post := &model.Post{}

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(fmt.Sprint(err)))
		return
	}

	post.Id = rand.Int()

	posts_repo.Save(post)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
