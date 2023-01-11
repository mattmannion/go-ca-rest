package mock_clients

import (
	"_/src/types/router_types"
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpClient struct {
	mock.Mock
}

func (MockHttpClient) ServeHttp(resp http.ResponseWriter, req *http.Request) {}

func (MockHttpClient) Mux() http.Handler {
	return http.NewServeMux()
}

func (MockHttpClient) Get(url string, f http.HandlerFunc) router_types.IRoute    { return f }
func (MockHttpClient) Post(url string, f http.HandlerFunc) router_types.IRoute   { return f }
func (MockHttpClient) Put(url string, f http.HandlerFunc) router_types.IRoute    { return f }
func (MockHttpClient) Patch(url string, f http.HandlerFunc) router_types.IRoute  { return f }
func (MockHttpClient) Delete(url string, f http.HandlerFunc) router_types.IRoute { return f }
