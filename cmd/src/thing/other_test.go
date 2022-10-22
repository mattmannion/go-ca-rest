package thing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.Equal(t, Hello(), "Hello", "should say \"Hello\"")
}

func TestAdd(t *testing.T) {
	assert.Equal(t, Add(1, 1), 2, "Add should equal 2")
	assert.NotEqual(t, Add(1, 1), 1, "Add should not equal 2")
}
