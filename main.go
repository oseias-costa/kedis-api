package main

import (
	// "encoding/json"
	"fmt"
	"main/infra/repository"
	"main/infra/router"
	"main/presentation/controller"
	"main/usecases"
	"net/http"
)

var httpRouter = router.NewMuxRouter()
var userUseCase = usecases.NewUserUseCase()
var userRepository = repository.NewUserRepository(userUseCase)
var userController = controller.NewUserController(userRepository)

func main() {
	fmt.Println("Hello, World!")
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Hello World"}`))
	})
	httpRouter.GET("/user/{id}", userController.GetUserByID)
	httpRouter.POST("/user/", userController.CreateNewUser)
	httpRouter.GET("/user/", userController.GetAllUsers)
	httpRouter.PUT("/user/", userController.UpdateUser)
	httpRouter.DELETE("/user/{id}", userController.DeleteUser)
	httpRouter.SERVE(":8100")
}
