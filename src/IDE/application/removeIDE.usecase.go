package application

import "rest/src/IDE/domain"

type RemoveIDEUseCase struct {
	repo domain.IDERepository
}

func NewRemoveIDEUseCase(repo domain.IDERepository) *RemoveIDEUseCase {
	return &RemoveIDEUseCase{repo: repo}
}

func (uc *RemoveIDEUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
