package gin_mux

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

type muxRouter struct{}

func NewMux() router_types.IMux {
	return &muxRouter{}
}

func (mr *muxRouter) Mux() http.Handler {
	return server
}

func gin_hf(f func(resp http.ResponseWriter, req *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) { f(c.Writer, c.Request) }
}

func (mr *muxRouter) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.GET(url, gin_hf(f))
}

func (*muxRouter) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.POST(url, gin_hf(f))
}
