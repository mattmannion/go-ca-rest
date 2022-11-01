package post_service

import (
	"_/src/api/repo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostService(t *testing.T) {
	assert.IsType(t, &PostService{}, NewPostService(repo.NewRepoLayer().PostRepo))
}
