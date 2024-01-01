package router

import "net/http"

type Router interface {
	SERVE(port string)
	GET(path string, handler func(http.ResponseWriter, *http.Request))
	POST(path string, handler func(http.ResponseWriter, *http.Request))
	PUT(path string, handler func(http.ResponseWriter, *http.Request))
	DELETE(path string, handler func(http.ResponseWriter, *http.Request))
}
