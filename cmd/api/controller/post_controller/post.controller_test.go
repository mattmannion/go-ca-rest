package post_controller

import (
	"_/cmd/api/service/post_service"
	"_/cmd/constants"
	"_/cmd/mocks/mock_services"
	"_/cmd/models"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	posts_route = constants.ApiPrefixV1 + "/posts"
	Validate    = "Validate"
	FindAll     = "FindAll"
	Create      = "Create"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostController{}, NewPostController(&post_service.PostService{}))
}

func TestGetPostsSuccess(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On(FindAll).Return([]models.Post{}, nil)

	PostController := NewPostController(MockService)

	req, err := http.NewRequest(http.MethodGet, posts_route, nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()

	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `[]`, rw.Body.String())
}

func TestGetPostsFailure(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On(FindAll).Return([]models.Post{}, errors.New("error"))

	PostController := NewPostController(MockService)

	req, err := http.NewRequest(http.MethodGet, posts_route, nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()

	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)

	assert.JSONEq(t, `{"error": "Could not find Posts"}`, rw.Body.String())
}

// func TestPostPostSuccess(t *testing.T) {
// 	Post := &models.Post{Id: 1, Title: "test", Text: "this"}

// 	MockService := new(mock_services.MockPostService)
// 	// MockService.On(Validate, Post).Return(nil)
// 	mc := MockService.On(Create, Post).Return(Post, nil)

// 	mc.RunFn = func(args mock.Arguments) {
// 		mc.ReturnArguments = args
// 	}

// 	PostController := NewPostController(MockService)

// 	body, err := json.Marshal(Post)

// 	assert.Equal(t, nil, err)

// 	req, err := http.NewRequest(http.MethodPost, posts_route, bytes.NewBuffer(body))

// 	assert.Equal(t, nil, err)

// 	rw := httptest.NewRecorder()

// 	http.HandlerFunc(PostController.PostPost).ServeHTTP(rw, req)

// 	assert.Equal(t, http.StatusOK, rw.Code)

// 	assert.JSONEq(t, `{"id": 1, "title": "test", "text": "this"}`, rw.Body.String())
// }

// func TestPostPostFailure(t *testing.T) {
// 	MockService := new(mock_services.MockPostService)
// 	MockService.On("").Return([]models.Post{}, errors.New("error"))

// 	PostController := NewPostController(MockService)

// 	req, err := http.NewRequest(http.MethodPost, posts_route, nil)

// 	assert.Equal(t, nil, err)

// 	rw := httptest.NewRecorder()

// 	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

// 	assert.Equal(t, http.StatusInternalServerError, rw.Code)

// 	assert.JSONEq(t, `{"error": "Could not find Posts"}`, rw.Body.String())
// }
