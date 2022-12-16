package muxes

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/gorilla/mux"
)

var gor_mux *mux.Router = mux.NewRouter()

type GorMux struct{}

func NewGorMux() router_types.IMux {
	return &GorMux{}
}

func (mr *GorMux) Mux() http.Handler {
	return gor_mux
}

func (*GorMux) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	gor_mux.HandleFunc(url, f).Methods(http.MethodGet)
}

func (*GorMux) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	gor_mux.HandleFunc(url, f).Methods(http.MethodPost)
}
