package usecases

import (
	"fmt"
	"main/domain"
	"main/infra/repository"
	"time"

	"github.com/google/uuid"
)

var repoResult repository.ResultRepository

type ResultUsecase interface {
	CreateResult(result domain.Result) (bool, error)
}

type resultUseCase struct{}

func NewResultUseCase(repo repository.ResultRepository) ResultUsecase {
	repoResult = repo
	return &resultUseCase{}
}

func (*resultUseCase) CreateResult(result domain.Result) (bool, error) {
	id, errUiid := uuid.NewRandom()
	if errUiid != nil {
		return false, errUiid
	}
	result.Id = id.String()
	now := time.Now()
	y, d, m := now.Date()
	result.Date = fmt.Sprintf("%d/%d/%d", d, int(m), y)

	err := repoResult.CreateResultRepo(result)
	if err != nil {
		return false, err
	}

	_, errWA := CreateWrogAnswer(result.Id, result.Details)
	if errWA != nil {
		return false, err
	}

	return true, nil
}

func CreateWrogAnswer(resultId string, wrongAnswers []domain.WrongAnswers) (bool, error) {
	for _, res := range wrongAnswers {
		idWrongAnswer, errWR := uuid.NewRandom()
		if errWR != nil {
			return false, errWR
		}
		res.Id = idWrongAnswer.String()
		res.ResultId = resultId

		err := repoResult.CreateWrongAnswerRepo(res)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
