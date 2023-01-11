package muxes

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var gin_mux *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)

	gin_mux = gin.New()

	gin_mux.Use(gin.Logger())
	gin_mux.Use(gin.Recovery())
}

type GinMux struct{}

func NewGinMux() router_types.IMux {
	return &GinMux{}
}

func (mr *GinMux) Mux() http.Handler {
	return gin_mux
}

func (*GinMux) Get(url string, f http.HandlerFunc) router_types.IRoute {
	return gin_mux.GET(url, gin.WrapF(f))
}

func (*GinMux) Post(url string, f http.HandlerFunc) router_types.IRoute {
	return gin_mux.POST(url, gin.WrapF(f))
}
