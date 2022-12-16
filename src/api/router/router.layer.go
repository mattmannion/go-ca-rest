package router

import (
	"_/src/api/controller"
	"_/src/api/router/routers"
	"_/src/envs"
	"_/src/types/router_types"
	"fmt"
	"log"
	"net/http"
)

type Deps struct {
	Mux   router_types.IMux
	Ctrlr controller.ControllerLayer
	Cfg   envs.TCfg
}

type RouterLayer struct {
	Mux     http.Handler
	Cfg     envs.TCfg
	Routers []router_types.IRouter
}

func NewRouterLayer(deps Deps) *RouterLayer {
	return &RouterLayer{
		Mux: deps.Mux.Mux(),
		Cfg: deps.Cfg,
		Routers: []router_types.IRouter{
			routers.NewPostRouter(deps.Mux, deps.Ctrlr),
		},
	}
}

func (rl *RouterLayer) ServeRestApi() {
	for _, router := range rl.Routers {
		router.Register()
	}

	fmt.Printf("Server live at: http://%s:%s\n", rl.Cfg.Host, rl.Cfg.Port)
	log.Fatalln(http.ListenAndServe(rl.Cfg.Host+":"+rl.Cfg.Port, rl.Mux))
}
