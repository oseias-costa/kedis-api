package main

import (
	"main/infra/initializers"
	"main/infra/router"
	"main/presentation/controller"
	"main/presentation/middlewares"
	"main/usecases"
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

	var examUseCase = usecases.NewExamUseCase()
	var examController = controller.NewExamController(examUseCase)

	httpRouter.GET("/user/", middlewares.Auth(userController.GetUser))
	httpRouter.POST("/user/", userController.CreateUser)
	httpRouter.POST("/login", userController.LoginUser)

	httpRouter.GET("/exam", examController.GetExam)

	httpRouter.SERVE(os.Getenv("PORT"))
}
