package routers

import (
	"_/src/api/controller"
	"_/src/mocks/mock_clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

var pr = NewPostRouter(mock_clients.MockHttpClient{}, controller.ControllerLayer{})

func TestNewPostRouter(t *testing.T) {
	assert.IsType(t, &PostRouter{}, pr)
}

func TestRegister(t *testing.T) {
	// pr.Register()
}
