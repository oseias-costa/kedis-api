package controller

import (
	"encoding/json"
	"fmt"
	"main/domain"
	"main/presentation/middlewares"
	"main/usecases"
	"net/http"
)

var examUseCase usecases.ExamUseCase

type ExamController interface {
	GetExam(w http.ResponseWriter, r *http.Request)
}

type examController struct{}

func NewExamController(u usecases.ExamUseCase) ExamController {
	examUseCase = u
	return &examController{}
}

func (*examController) GetExam(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var exam domain.ExamReq
	id := middlewares.GetUserId(w, r)

	err := json.NewDecoder(r.Body).Decode(&exam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "user not decoded"}`))
	}

	fmt.Println("esse é o id do auto", id)
	fmt.Println("esse é o body do Ver se deu", exam)

	arr, err := examUseCase.GetExamUseCase(fmt.Sprintf("./assets/%s.json", exam.ExamName))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(arr))
}
