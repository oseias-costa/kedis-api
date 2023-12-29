package main

import (
	// "encoding/json"
	"fmt"
	"main/infra/repository"
	"main/infra/router"
	"net/http"
)

var httpRouter = router.NewMuxRouter()
var userRoutes = repository.NewUserRepository()

func main() {
	fmt.Println("Hello, World!")
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Hello World"}`))
	})
	httpRouter.GET("/user/{id}", userRoutes.GetUserByID)
	httpRouter.SERVE(":8100")
}
