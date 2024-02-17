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
	CreateResult(result domain.Result) (string, error)
	GetResultById(resultId string) (domain.Result, error)
}

type resultUseCase struct{}

func NewResultUseCase(repo repository.ResultRepository) ResultUsecase {
	repoResult = repo
	return &resultUseCase{}
}

func (*resultUseCase) CreateResult(result domain.Result) (string, error) {
	id, errUiid := uuid.NewRandom()
	if errUiid != nil {
		return "", errUiid
	}
	result.Id = id.String()
	now := time.Now()
	y, m, d := now.Date()
	result.Date = fmt.Sprintf("%d/%d/%d", d, int(m), y)

	err := repoResult.CreateResultRepo(result)
	if err != nil {
		return "", err
	}

	_, errWA := CreateWrogAnswer(result.Id, result.Details)
	if errWA != nil {
		return "", err
	}

	resultId := fmt.Sprintf(`{"resultId": %s}`, id)

	return resultId, nil
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

func (*resultUseCase) GetResultById(resultId string) (domain.Result, error) {
	r, err := repoResult.GetResultRepo(resultId)
	if err != nil {
		return r, err
	}

	return r, err
}
