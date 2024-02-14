package controller

import (
	"encoding/json"
	"fmt"
	"main/domain"
	"main/presentation/middlewares"
	"main/usecases"
	"net/http"
)

var resultUseCase usecases.ResultUsecase

type ResultController interface {
	CreateResults(w http.ResponseWriter, r *http.Request)
}

type resultController struct{}

func NewResultController(usecase usecases.ResultUsecase) ResultController {
	resultUseCase = usecase
	return &resultController{}
}

func (*resultController) CreateResults(w http.ResponseWriter, r *http.Request) {
	id := middlewares.GetUserId(w, r)
	var result domain.Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "result not decoded"}`))
	}
	result.UserId = id
	fmt.Println("esse é o id do auto", id)

	u, err := resultUseCase.CreateResult(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprint(err)))
	}
	fmt.Println("return controller CreateResults", u)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
