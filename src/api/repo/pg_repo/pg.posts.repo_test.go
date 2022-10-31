package pg_repo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo())
}
