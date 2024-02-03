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
	muxDispatcher.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			// Continue with the request
			next.ServeHTTP(w, r)
		})
	})
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
