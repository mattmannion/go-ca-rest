package firestore_repo

import (
	"_/src/envs"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
)

func TestNewPostRepo(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(firestore.NewClient, envs.FirestoreProd))
}

func TestSaveClientFailure(t *testing.T) {

}
