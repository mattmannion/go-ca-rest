package mock_controllers

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockPostController struct {
	mock.Mock
}

func (c MockPostController) GetPosts(resp http.ResponseWriter, req *http.Request) {}

func (c MockPostController) PostPost(resp http.ResponseWriter, req *http.Request) {}
