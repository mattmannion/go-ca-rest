package router

import (
	"_/src/api/controller"
	"_/src/envs"
	"_/src/mocks/mock_clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rl = NewRouterLayer(
	Deps{
		Router:     mock_clients.MockHttpClient{},
		Controller: controller.ControllerLayer{},
		Cfg:        envs.Cfg,
	},
)

func TestNewRouterLayer(t *testing.T) {
	assert.IsType(t, &RouterLayer{}, rl)
}

func TestServeRestApi(t *testing.T) {
	// err := rl.ServeRestApi()
	// assert.Equal(t, nil, err)
}
