package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonAwsExam struct {
	Id                         int    `json:"id"`
	ServiceType                string `json:"serviceType"`
	Topic                      string `json:"topic"`
	Question                   string `json:"question"`
	Opcoes                     string `json:"opcoes"`
	CorrectAlternative         string `json:"correctAlternative"`
	CorrectAlternativeFeedback string `json:"correctAlternativeFeedback"`
}

func ExameAwsRepository(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	var obj []JsonAwsExam

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println(err)
	}

	return string(data), err
}
