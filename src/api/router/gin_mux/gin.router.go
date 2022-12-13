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

func NewMuxRouter() router_types.IRouter {
	return &muxRouter{}
}

func (mr *muxRouter) Handler() http.Handler {
	return server
}

func gin_wrapper(f func(resp http.ResponseWriter, req *http.Request)) gin.HandlerFunc {
	return func(c *gin.Context) { f(c.Writer, c.Request) }
}

func (mr *muxRouter) Get(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.GET(url, gin_wrapper(f))
}

func (*muxRouter) Post(url string, f func(resp http.ResponseWriter, req *http.Request)) {
	server.POST(url, gin_wrapper(f))
}
