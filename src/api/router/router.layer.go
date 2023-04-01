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
	Cfg        envs.TCfg
}

type RouterLayer struct {
	Mux     http.Handler
	Cfg     envs.TCfg
	Routers []router_types.IRouter
}

func NewRouterLayer(deps Deps) *RouterLayer {
	return &RouterLayer{
		Mux: deps.Router.Mux(),
		Cfg: deps.Cfg,
		Routers: []router_types.IRouter{
			routers.NewPostRouter(deps.Router, deps.Controller),
		},
	}
}

func (rl *RouterLayer) ServeRestApi() error {
	for _, router := range rl.Routers {
		router.Register()
	}

	fmt.Printf("Server live at: http://%s:%s\n", rl.Cfg.Host, rl.Cfg.Port)
	return http.ListenAndServe(rl.Cfg.Host+":"+rl.Cfg.Port, rl.Mux)
}
