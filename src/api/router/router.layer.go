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
	handler    http.Handler
	cfg        envs.TCfg
	PostRouter routers.PostRouter
}

func NewRouterLayer(
	router router_types.IRouter,
	ctrlr controller.ControllerLayer,
	cfg envs.TCfg,
) *RouterLayer {
	return &RouterLayer{
		handler:    router.Handler(),
		cfg:        cfg,
		PostRouter: *routers.NewPostRouter(router, ctrlr)}
}

func (rl *RouterLayer) Serve() {
	rl.PostRouter.Register()

	fmt.Printf("Server live at: http://%s:%s\n", rl.cfg.Host, rl.cfg.Port)
	log.Fatalln(http.ListenAndServe(rl.cfg.Host+":"+rl.cfg.Port, rl.handler))
}
