package gmux

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/gorilla/mux"
)

var gmux_router *mux.Router = mux.NewRouter()

type muxRouter struct{}

func NewMux() router_types.IMux {
	return &muxRouter{}
}

func (mr *muxRouter) Mux() http.Handler {
	return gmux_router
}

func (*muxRouter) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	gmux_router.HandleFunc(url, f).Methods(http.MethodGet)
}

func (*muxRouter) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	gmux_router.HandleFunc(url, f).Methods(http.MethodPost)
}
