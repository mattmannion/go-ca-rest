package router

import (
	"_/src/api/controller"
	"_/src/api/router/routers"
	"_/src/envs"
	"_/src/types/router_types"
	"fmt"
	"net/http"
)

type Deps struct {
	Router     router_types.IMux
	Controller controller.ControllerLayer
}

type RouterLayer struct {
	Mux     http.Handler
	Routers []router_types.IRouter
}

func NewRouterLayer(deps Deps) *RouterLayer {
	return &RouterLayer{
		Mux: deps.Router.Mux(),
		Routers: []router_types.IRouter{
			routers.NewPostRouter(deps.Router, deps.Controller),
		},
	}
}

func (rl *RouterLayer) CreateRestApi() (string, http.Handler) {
	cfg := envs.Cfg

	for _, router := range rl.Routers {
		router.Register()
	}

	fmt.Printf("Server live at: http://%s:%s\n", cfg.Host, cfg.Port)
	return cfg.Host + ":" + cfg.Port, rl.Mux
}
