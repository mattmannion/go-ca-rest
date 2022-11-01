package pg_repo

import (
	"_/src/clients/pg/pg_client"
	"_/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(pg_client.Db))
}

func TestInsert(t *testing.T) {
	PostRepo := NewPostRepo(pg_client.Db)

	Post := &models.Post{
		Title: "Test Post",
		Text:  "Test Post",
	}

	result, err := PostRepo.Insert(Post)

	assert.Equal(t, nil, err)

	assert.IsType(t, &models.Post{
		Id:    4,
		Title: Post.Title,
		Text:  Post.Text,
	}, result)

	pg_client.ResetAndSeedPgDb()
}

func TestGetAll(t *testing.T) {
	PostRepo := NewPostRepo(pg_client.Db)

	Posts, err := PostRepo.GetAll()

	assert.Equal(t, nil, err)

	assert.IsType(t, []models.Post{}, Posts)

	pg_client.ResetAndSeedPgDb()
}
