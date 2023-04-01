package pg_repo

import (
	"_/src/clients/pg"
	"_/src/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostRepo{}, NewPostRepo(pg.Db))
}

func TestInsert(t *testing.T) {
	PostRepo := NewPostRepo(pg.Db)

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

	pg.ResetAndSeedPgDb()
}

func TestGetAll(t *testing.T) {
	PostRepo := NewPostRepo(pg.Db)

	Posts, err := PostRepo.GetAll()

	assert.Equal(t, nil, err)

	assert.IsType(t, []models.Post{}, Posts)

	pg.ResetAndSeedPgDb()
}
