package repository

import (
	"fmt"
	"main/domain"
	"main/infra/persistence"
)

type ResultRepository interface {
	CreateResultRepo(r domain.Result) error
	CreateWrongAnswerRepo(r domain.WrongAnswers) error
	GetResultRepo(resultId string) (domain.Result, error)
}

type resultRepo struct{}

func NewResultRepo() ResultRepository {
	return &resultRepo{}
}

func (*resultRepo) CreateResultRepo(r domain.Result) error {
	c := persistence.Connect()

	sql := "INSERT INTO results (id, userId, date, cloud, mockExam, result, wrong, correct) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	stmt, err := c.Prepare(sql)
	if err != nil {
		return err
	}
	//stmt.Close()

	res, err := stmt.Exec(r.Id, r.UserId, r.Date, r.Cloud, r.MockExam, r.Result, r.Wrong, r.Correct)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func (*resultRepo) CreateWrongAnswerRepo(w domain.WrongAnswers) error {
	c := persistence.Connect()

	sql := "INSERT INTO wrongAnswers (id, resultId, serviceType, topic, wrongAnswers, total) VALUES (?, ?, ?, ?, ?, ?)"

	stmt, err := c.Prepare(sql)
	if err != nil {
		return err
	}
	//stmt.Close()

	res, err := stmt.Exec(w.Id, w.ResultId, w.ServiceType, w.Topic, w.WrongsAnswers, w.Total)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func (*resultRepo) GetResultRepo(resultId string) (domain.Result, error) {
	var result domain.Result
	c := persistence.Connect()

	r, err := c.Query("SELECT * FROM results WHERE id = ?", resultId)
	if err != nil {
		return result, err
	}

	for r.Next() {
		err := r.Scan(
			&result.Id,
			&result.UserId,
			&result.Date,
			&result.Cloud,
			&result.Result,
			&result.Wrong,
			&result.Correct,
			&result.MockExam,
		)
		if err != nil {
			return result, err
		}
	}
	defer c.Close()
	defer r.Close()

	return result, nil
}
