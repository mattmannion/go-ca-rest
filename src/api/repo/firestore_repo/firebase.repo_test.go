package firestore_repo

import (
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
)

func TestNewPostRepo(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(firestore.NewClient))
}

func TestSaveClientFailure(t *testing.T) {
}
