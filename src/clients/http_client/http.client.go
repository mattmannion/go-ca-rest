package http_client

import (
	"log"
	"net/http"
)

func ListenAndServeClient(addr string, handler http.Handler) {
	log.Fatalln(http.ListenAndServe(addr, handler))
}
