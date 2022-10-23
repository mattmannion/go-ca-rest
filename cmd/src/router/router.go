package router

import "net/http"

type Router interface {
	Get(url string, f func(resp http.ResponseWriter, req *http.Request))
	Post(url string, f func(resp http.ResponseWriter, req *http.Request))
	Serve(port string)
}
