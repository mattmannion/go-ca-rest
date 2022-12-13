package gmux

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/gorilla/mux"
)

var mux_router *mux.Router = mux.NewRouter()

type muxRouter struct{}

func NewMuxRouter() router_types.IRouter {
	return &muxRouter{}
}

func (mr *muxRouter) Handler() http.Handler {
	return mux_router
}

func (*muxRouter) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	mux_router.HandleFunc(url, f).Methods(http.MethodGet)
}

func (*muxRouter) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	mux_router.HandleFunc(url, f).Methods(http.MethodPost)
}
