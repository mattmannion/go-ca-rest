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
	Router router_types.IMux
	Ctrlr  controller.ControllerLayer
	Cfg    envs.TCfg
}

type RouterLayer struct {
	handler http.Handler
	cfg     envs.TCfg
	routers []router_types.IRouter
}

func NewRouterLayer(deps Deps) *RouterLayer {
	return &RouterLayer{
		handler: deps.Router.Handler(),
		cfg:     deps.Cfg,
		routers: []router_types.IRouter{
			routers.NewPostRouter(deps.Router, deps.Ctrlr),
		},
	}
}

func (rl *RouterLayer) ServeRestApi() {
	for _, router := range rl.routers {
		router.Register()
	}

	fmt.Printf("Server live at: http://%s:%s\n", rl.cfg.Host, rl.cfg.Port)
	log.Fatalln(http.ListenAndServe(rl.cfg.Host+":"+rl.cfg.Port, rl.handler))
}
