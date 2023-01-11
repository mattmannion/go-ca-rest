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

func TestGorGet(t *testing.T) {
	mux := NewGorMux()
	assert.IsType(t, &gm.Route{}, mux.Get("/gortesting", fn))
}

func TestGorPost(t *testing.T) {
	mux := NewGorMux()
	assert.IsType(t, &gm.Route{}, mux.Post("/gortesting", fn))
}
