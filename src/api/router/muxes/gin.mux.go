package muxes

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)

	server = gin.New()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
}

type GinMux struct{}

func NewGinMux() router_types.IMux {
	return &GinMux{}
}

func (mr *GinMux) Mux() http.Handler {
	return server
}

func gin_hf(f func(resp http.ResponseWriter, req *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) { f(c.Writer, c.Request) }
}

func (mr *GinMux) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.GET(url, gin_hf(f))
}

func (*GinMux) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.POST(url, gin_hf(f))
}
