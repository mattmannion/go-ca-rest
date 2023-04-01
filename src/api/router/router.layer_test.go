package router

import (
	"_/src/api/controller"
	"_/src/mocks/mock_clients"
	"_/src/mocks/mock_controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rl = NewRouterLayer(
	Deps{
		Router:     mock_clients.MockHttpClient{},
		Controller: controller.ControllerLayer{PostController: mock_controllers.MockPostController{}},
	},
)

func TestNewRouterLayer(t *testing.T) {
	assert.IsType(t, RouterLayer{}, *rl)
}

func TestCreateRestApi(t *testing.T) {
	rl.CreateRestApi()
}
