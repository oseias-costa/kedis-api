package usecases

import "main/domain"

type ResultUsecase interface {
	CreateResult(result domain.Result)
}

type resultUseCase struct{}

func NewResultUseCase() ResultUsecase {
	return &resultUseCase{}
}

func (*resultUseCase) CreateResult(result domain.Result) {

}
