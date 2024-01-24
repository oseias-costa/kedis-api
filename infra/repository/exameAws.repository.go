package repository

import (
	"fmt"
	"os"
)

type JsonAwsExam struct {
	Id                         int    `json:"id"`
	ServiceType                string `json:"serviceType"`
	Topic                      string `json:"topic"`
	Question                   string `json:"question"`
	Opcoes                     string `json:"answers"`
	CorrectAlternative         string `json:"correctAlternative"`
	CorrectAlternativeFeedback string `json:"correctAlternativeFeedback"`
}

type ExamRepository interface {
	GetExamRepo(path string) (string, error)
}

type examRepository struct{}

func NewExamRepository() ExamRepository {
	return &examRepository{}
}

func (*examRepository) GetExamRepo(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	// var obj []JsonAwsExam

	// err = json.Unmarshal(data, &obj)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return string(data), err
}
