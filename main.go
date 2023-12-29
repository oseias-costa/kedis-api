package main

import (
	"fmt"
	"main/infra/router"
	"main/usecases"
)

var httpRouter = router.NewMuxRouter()
var useCase = usecases.NewUserUseCase()

func main() {
	fmt.Println("Hello, World!")
  httpRouter.GET("/user/{id}", useCase.GetUser)
	httpRouter.SERVE(":8100")
}
