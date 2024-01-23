package usecases

import (
	"main/domain"
	"main/infra/repository"

	"github.com/google/uuid"
)

var repoResult = repository.NewResultRepo()

type ResultUsecase interface {
	CreateResult(result domain.Result) (bool, error)
}

type resultUseCase struct{}

func NewResultUseCase() ResultUsecase {
	return &resultUseCase{}
}

func (*resultUseCase) CreateResult(result domain.Result) (bool, error) {
	id, errUiid := uuid.NewRandom()
	if errUiid != nil {
		return false, errUiid
	}
	result.Id = id.String()

	err := repoResult.CreateResultRepo(result)
	if err != nil {
		return false, err
	}

	for _, res := range result.Details {
		idWrongAnswer, errWR := uuid.NewRandom()
		if errWR != nil {
			return false, errWR
		}
		res.Id = idWrongAnswer.String()
		res.ResultId = result.Id

		err := repoResult.CreateWrongAnswerRepo(res)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}
