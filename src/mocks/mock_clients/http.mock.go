package mock_clients

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpClient struct {
	mock.Mock
}

func (MockHttpClient) ServeHttp(resp http.ResponseWriter, req *http.Request) {

}
