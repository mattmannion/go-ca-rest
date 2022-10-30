package firestore_repo

import (
	"_/src/mocks/mock_clients"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	mock_clients.TestMain(m)
}

func TestNewPostRepo(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(mock_clients.NewFirestoreTestClient))
}

func TestSaveClientFailure(t *testing.T) {
}
