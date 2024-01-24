package controller

import (
	"fmt"
	"io"
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
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}

	fmt.Println("esse é o body do GetExam", b)

	arr, err := examUseCase.GetExamUseCase("./assets/cloud-practictioner#1.json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(arr))
}
