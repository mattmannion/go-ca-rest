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

type RouterLayer struct {
	handler http.Handler
	cfg     envs.TCfg
	routers []router_types.IRouter
}

func NewRouterLayer(
	router router_types.IMux,
	ctrlr controller.ControllerLayer,
	cfg envs.TCfg,
) *RouterLayer {
	return &RouterLayer{
		handler: router.Handler(),
		cfg:     cfg,
		routers: []router_types.IRouter{
			routers.NewPostRouter(router, ctrlr),
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
