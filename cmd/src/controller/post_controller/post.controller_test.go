package post_controller

import (
	"_/cmd/mocks/mock_services"
	"_/cmd/models"
	"_/cmd/src/service/post_service"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPostController(t *testing.T) {
	assert.IsType(t, &PostController{}, NewPostController(&post_service.PostService{}))
}

func TestGetPostsSuccess(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On("FindAll").Return([]models.Post{}, nil)

	PostController := NewPostController(MockService)

	req, err := http.NewRequest("GET", "/", nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()
	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `[]`, rw.Body.String())
}

func TestGetPostsFailure(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On("FindAll").Return([]models.Post{}, errors.New("error"))

	PostController := NewPostController(MockService)

	req, err := http.NewRequest("GET", "/", nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()
	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)

	assert.JSONEq(t, `{"error": "Could not find Posts..."}`, rw.Body.String())
}

func TestPostPostSuccess(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On("FindAll").Return([]models.Post{}, nil)

	PostController := NewPostController(MockService)

	req, err := http.NewRequest("GET", "/", nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()
	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusOK, rw.Code)

	assert.JSONEq(t, `[]`, rw.Body.String())
}

func TestPostPostFailure(t *testing.T) {
	MockService := new(mock_services.MockPostService)
	MockService.On("FindAll").Return([]models.Post{}, errors.New("error"))

	PostController := NewPostController(MockService)

	req, err := http.NewRequest("GET", "/", nil)

	assert.Equal(t, nil, err)

	rw := httptest.NewRecorder()
	http.HandlerFunc(PostController.GetPosts).ServeHTTP(rw, req)

	assert.Equal(t, http.StatusInternalServerError, rw.Code)

	assert.JSONEq(t, `{"error": "Could not find Posts..."}`, rw.Body.String())
}
