package routers

import (
	"_/src/api/controller"
	"_/src/mocks/mock_clients"
	"_/src/mocks/mock_controllers"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pr = NewPostRouter(mock_clients.MockHttpClient{}, controller.ControllerLayer{PostController: mock_controllers.MockPostController{}})

func TestNewPostRouter(t *testing.T) {
	assert.IsType(t, &PostRouter{}, pr)
}

func TestRegister(t *testing.T) {
	pr.Register()
}
