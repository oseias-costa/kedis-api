package controller

import (
	"encoding/json"
	"main/domain"
	"main/presentation/middlewares"
	"main/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

var resultUseCase usecases.ResultUsecase

type ResultController interface {
	CreateResults(w http.ResponseWriter, r *http.Request)
	GetResultById(w http.ResponseWriter, r *http.Request)
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

	u, err := resultUseCase.CreateResult(result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (*resultController) GetResultById(w http.ResponseWriter, r *http.Request) {
	resultId := mux.Vars(r)["resultId"]
	if resultId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "resultId invalid"}`))
	}

	result, err := resultUseCase.GetResultById(resultId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
