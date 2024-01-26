package main

import (
	"main/infra/initializers"
	"main/infra/repository"
	"main/infra/router"
	"main/presentation/controller"
	"main/presentation/middlewares"
	"main/usecases"
	"os"
)

var httpRouter = router.NewMuxRouter()
var userController = controller.NewUserController()

var resultRepo = repository.NewResultRepo()
var resultUseCase = usecases.NewResultUseCase(resultRepo)
var resultController = controller.NewResultController(resultUseCase)

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

	httpRouter.POST("/result", middlewares.Auth(resultController.CreateResults))
	httpRouter.GET("/exam", middlewares.Auth(examController.GetExam))

	httpRouter.SERVE(os.Getenv("PORT"))
}
