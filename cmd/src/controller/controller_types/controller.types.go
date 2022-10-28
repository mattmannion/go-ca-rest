package controller_types

import "net/http"

type IPostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	PostPost(resp http.ResponseWriter, req *http.Request)
}
