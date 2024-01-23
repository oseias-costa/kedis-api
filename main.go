package main

import (
	"fmt"
	"main/infra/initializers"
	"main/infra/repository"
	"main/infra/router"
	"main/presentation/controller"
	"main/presentation/middlewares"
	"os"
)

var httpRouter = router.NewMuxRouter()
var userController = controller.NewUserController()

type Claims struct {
	Id string
}

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	// users, err := persistence.GetAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	data, err := repository.ExameAwsRepository("./assets/cloud-practictioner#1.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
	// fmt.Println("user list")
	// fmt.Println(users)

	httpRouter.GET("/user/", middlewares.Auth(userController.GetUser))
	httpRouter.POST("/user/", userController.CreateUser)
	httpRouter.POST("/login", userController.LoginUser)

	httpRouter.SERVE(os.Getenv("PORT"))
}
