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
	assert.IsType(t, &gm.Route{}, mux.Get("/gortesting", fn))
	assert.IsType(t, &gm.Route{}, mux.Post("/gortesting", fn))
	assert.IsType(t, &gm.Route{}, mux.Put("/gortesting", fn))
	assert.IsType(t, &gm.Route{}, mux.Patch("/gortesting", fn))
	assert.IsType(t, &gm.Route{}, mux.Delete("/gortesting", fn))
}
