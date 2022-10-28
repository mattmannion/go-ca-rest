package controller

import (
	"_/cmd/src/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewControllerLayer(t *testing.T) {
	assert.IsType(t, &ControllerLayer{}, NewControllerLayer(Deps{ServiceLayer: service.ServiceLayer{}}))
}
