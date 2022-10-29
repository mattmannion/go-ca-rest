package post_controller

import (
	"_/cmd/api/service/post_service"
	"_/cmd/constants"
	"_/cmd/mocks/mock_services"
	"_/cmd/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	posts_route string = constants.ApiPrefixV1 + "/posts"
	Validate    string = "Validate"
	FindAll     string = "FindAll"
	Create      string = "Create"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostController{}, NewPostController(&post_service.PostService{}))
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

func TestGetPostsSuccess(t *testing.T) {
	// checks the endpoint twice to ensure consistency
	MockService := new(mock_services.MockPostService)
	MockService.On(FindAll).Return([]models.Post{}, nil)

	PostController := NewPostController(MockService)

	// first call
	req, err := http.NewRequest(http.MethodGet, posts_route, nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()

	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `[]`, rw.Body.String())

	// second call
	req, err = http.NewRequest(http.MethodGet, posts_route, nil)

	assert.Equal(t, nil, err)

	rw = httptest.NewRecorder()

	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `[]`, rw.Body.String())
}

// func TestPostPostValidationFailure(t *testing.T) {
// 	Post := &models.Post{}

// 	MockService := new(mock_services.MockPostService)
// 	MockService.On(Validate, mock.Anything).Return(errors.New("errors"))
// 	MockService.On(Create, mock.Anything).Return(Post, nil)

// 	PostController := NewPostController(MockService)

// 	body, err := json.Marshal(Post)

// 	assert.Equal(t, nil, err)

// 	req, err := http.NewRequest(http.MethodPost, posts_route, bytes.NewBuffer(body))

// 	assert.Equal(t, nil, err)

// 	rw := httptest.NewRecorder()

// 	http.HandlerFunc(PostController.PostPost).ServeHTTP(rw, req)

// 	assert.Equal(t, http.StatusInternalServerError, rw.Code)

// 	assert.JSONEq(t, `{"id": 0, "title": "", "text": ""}`, rw.Body.String())

// }

func TestPostPostFailures(t *testing.T) {
	Post := &models.Post{}

	// validation errors
	MockService := new(mock_services.MockPostService)
	MockService.On(Validate, mock.Anything).Return(errors.New("post validation error"))
	MockService.On(Create, mock.Anything).Return(Post, nil)

	PostController := NewPostController(MockService)

	body, err := json.Marshal(Post)

	assert.Equal(t, nil, err)

	req, err := http.NewRequest(http.MethodPost, posts_route, bytes.NewBuffer(body))

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()

	http.HandlerFunc(PostController.PostPost).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)

	assert.JSONEq(t, `{"error": "post validation error"}`, rw.Body.String())

	// creation errors
	MockService = new(mock_services.MockPostService)
	MockService.On(Validate, mock.Anything).Return(nil)
	MockService.On(Create, mock.Anything).Return(Post, errors.New("post creation error"))

	PostController = NewPostController(MockService)

	body, err = json.Marshal(Post)

	assert.Equal(t, nil, err)

	req, err = http.NewRequest(http.MethodPost, posts_route, bytes.NewBuffer(body))

	assert.Equal(t, nil, err)

	rw = httptest.NewRecorder()

	http.HandlerFunc(PostController.PostPost).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)

	assert.JSONEq(t, `{"error": "post creation error"}`, rw.Body.String())

}

func TestPostPostSuccess(t *testing.T) {
	Post := &models.Post{Id: 1, Title: "test", Text: "this"}

	MockService := new(mock_services.MockPostService)
	MockService.On(Validate, mock.Anything).Return(nil)
	MockService.On(Create, mock.Anything).Return(Post, nil)

	PostController := NewPostController(MockService)

	body, err := json.Marshal(Post)

	assert.Equal(t, nil, err)

	req, err := http.NewRequest(http.MethodPost, posts_route, bytes.NewBuffer(body))

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()

	http.HandlerFunc(PostController.PostPost).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `{"id": 1, "title": "test", "text": "this"}`, rw.Body.String())
}
