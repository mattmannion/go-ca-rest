package post_service

import (
	"_/src/api/repo"
	mock_repo "_/src/mocks/mock_repos"
	"_/src/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetAll string = "GetAll"
	Insert string = "Insert"
)

func TestNewPostService(t *testing.T) {
	assert.IsType(t, &PostService{}, NewPostService(repo.NewRepoLayer().PostRepo))
}

func TestValidatePost(t *testing.T) {
	MockService := NewPostService(&mock_repo.MockPostRepo{})

	// no title
	post := &models.Post{}
	assert.Equal(t, errors.New("no title"), MockService.Validate(post))

	// no text
	post = &models.Post{Title: "a title"}
	assert.Equal(t, errors.New("no text"), MockService.Validate(post))

	// no errors
	post = &models.Post{Title: "a title", Text: "some text"}
	assert.Equal(t, nil, MockService.Validate(post))
}

func TestFindAllPosts(t *testing.T) {
	// test for failure
	MockRepo := new(mock_repo.MockPostRepo)
	MockRepo.On(GetAll).Return([]models.Post{}, errors.New("error"))

	MockService := NewPostService(MockRepo)

	_, err := MockService.FindAll()
	assert.Equal(t, errors.New("error"), err)

	// test for success
	MockRepo = new(mock_repo.MockPostRepo)
	MockRepo.On(GetAll).Return([]models.Post{}, nil)

	MockService = NewPostService(MockRepo)

	posts, err := MockService.FindAll()
	assert.Equal(t, nil, err)
	assert.Equal(t, []models.Post{}, posts)
}

func TestCreatePost(t *testing.T) {
	new_post := &models.Post{}

	// test for failure
	MockRepo := new(mock_repo.MockPostRepo)
	MockRepo.On(Insert, mock.Anything).Return(new_post, errors.New("error"))

	MockService := NewPostService(MockRepo)

	_, err := MockService.Create(new_post)
	assert.Equal(t, errors.New("error"), err)

	// test for success
	MockRepo = new(mock_repo.MockPostRepo)
	MockRepo.On(Insert, mock.Anything).Return(new_post, nil)

	MockService = NewPostService(MockRepo)

	posts, err := MockService.Create(new_post)
	assert.Equal(t, nil, err)
	assert.Equal(t, &models.Post{}, posts)
}
