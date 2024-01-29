package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var muxDispatcher = mux.NewRouter()

type muxRouter struct{}

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Listen and Serve on port %s\n", port)
	http.ListenAndServe(port, muxDispatcher)
}

func (*muxRouter) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(path, handler).Methods(http.MethodGet, http.MethodOptions)
}

func (*muxRouter) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(path, handler).Methods(http.MethodPost, http.MethodOptions)
}

func (*muxRouter) PUT(path string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(path, handler).Methods(http.MethodPut, http.MethodOptions)
}

func (*muxRouter) DELETE(path string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(path, handler).Methods(http.MethodDelete, http.MethodOptions)
}
