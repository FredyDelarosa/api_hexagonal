package application

import (
	"rest/src/ide/domain"
	"rest/src/ide/domain/entities"
)

type UpdateIDEUseCase struct {
	repo domain.IDERepository
}

func NewUpdateIDEUseCase(repo domain.IDERepository) *UpdateIDEUseCase {
	return &UpdateIDEUseCase{repo: repo}
}

func (uc *UpdateIDEUseCase) Execute(ide *entities.IDE) error {
	return uc.repo.Update(ide)
}
