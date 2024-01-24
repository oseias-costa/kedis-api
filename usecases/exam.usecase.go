package usecases

import "main/infra/repository"

var examRepo = repository.NewExamRepository()

type ExamUseCase interface {
	GetExamUseCase(e string) (string, error)
}

type examUseCase struct{}

func NewExamUseCase() ExamUseCase {
	return &examUseCase{}
}

func (*examUseCase) GetExamUseCase(e string) (string, error) {
	data, err := examRepo.GetExamRepo(e)
	if err != nil {
		return "", err
	}

	return data, nil
}
