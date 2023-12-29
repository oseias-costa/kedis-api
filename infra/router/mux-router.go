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
	http.ListenAndServe(port, nil)
}

func (*muxRouter) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	muxDispatcher.HandleFunc(path, handler).Methods(http.MethodGet)
}
