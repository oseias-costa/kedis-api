package main

import (
	"fmt"
	"main/infra/initializers"
	"main/infra/persistence"
	"main/infra/router"
	"main/presentation/controller"
	"net/http"
	"os"
)

var httpRouter = router.NewMuxRouter()
var userController = controller.NewUserController()

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello, World!")

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Hello World"}`))
	})

	users, err := persistence.GetAll()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("user list")
	fmt.Println(users)

	// httpRouter.GET("/user/{id}", userController.GetUserByID)
	httpRouter.POST("/user/", userController.CreateUser)
	httpRouter.SERVE(os.Getenv("PORT"))
}
