package muxes

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var fn http.HandlerFunc = func(resp http.ResponseWriter, req *http.Request) {}
var eng *gin.Engine

func TestNewGinMux(t *testing.T) {
	assert.IsType(t, &GinMux{}, NewGinMux())
}

func TestGinMux(t *testing.T) {
	mux := NewGinMux()
	assert.IsType(t, &gin.Engine{}, mux.Mux())
}

func TestGinMethods(t *testing.T) {
	mux := NewGinMux()
	assert.IsType(t, eng, mux.Get("", fn))
	assert.IsType(t, eng, mux.Post("", fn))
	assert.IsType(t, eng, mux.Put("", fn))
	assert.IsType(t, eng, mux.Patch("", fn))
	assert.IsType(t, eng, mux.Delete("", fn))
}
