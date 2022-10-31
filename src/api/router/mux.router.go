package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var mux_router *mux.Router = mux.NewRouter()

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	mux_router.HandleFunc(url, f).Methods(http.MethodGet)
}

func (*muxRouter) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	mux_router.HandleFunc(url, f).Methods(http.MethodPost)
}

func (*muxRouter) Serve(host string, port string) {
	fmt.Printf("Server live at: http://%s:%s\n", host, port)
	log.Fatalln(http.ListenAndServe(host+":"+port, mux_router))
}
