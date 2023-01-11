package muxes

import (
	"testing"

	gm "github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestNewGorMux(t *testing.T) {
	assert.IsType(t, &GorMux{}, NewGorMux())
}

func TestGorMux(t *testing.T) {
	mux := NewGorMux()
	assert.IsType(t, &gm.Router{}, mux.Mux())
}

func TestGorMethods(t *testing.T) {
	mux := NewGorMux()
	assert.IsType(t, &gm.Route{}, mux.Get("", fn))
	assert.IsType(t, &gm.Route{}, mux.Post("", fn))
	assert.IsType(t, &gm.Route{}, mux.Put("", fn))
	assert.IsType(t, &gm.Route{}, mux.Patch("", fn))
	assert.IsType(t, &gm.Route{}, mux.Delete("", fn))
}
