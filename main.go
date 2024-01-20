package main

import (
	"fmt"
	"main/infra/initializers"
	"main/infra/router"
	"main/presentation/controller"
	"main/presentation/middlewares"
	"net/http"
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
	// fmt.Println("Hello, World!")

	// httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(`{"message": "Hello World"}`))
	// })

	// users, err := persistence.GetAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("user list")
	// fmt.Println(users)

	httpRouter.GET("/user/", middlewares.Auth(userController.GetUser))
	httpRouter.POST("/user/", userController.CreateUser)
	httpRouter.POST("/login", userController.LoginUser)

	// r := mux.NewRouter()

	// r.HandleFunc("/login", userController.LoginUser).Methods(http.MethodPost)
	// r.HandleFunc("/", middlewares.Auth(Test)).Methods(http.MethodGet)

	httpRouter.SERVE(os.Getenv("PORT"))
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello word")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello word"))
}
