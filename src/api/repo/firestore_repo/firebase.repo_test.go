package firestore_repo

import (
	"_/src/envs"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
)

func TestNewPostRepo(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(firestore.NewClient, envs.FirestoreProjectName, envs.FirestoreCollectionName))
}

func TestSaveClientFailure(t *testing.T) {
	// PostRepo := NewPostRepo("", envs.FirestoreCollectionName)

	// assert.Equal(t, "err", PostRepo.Save())
}
