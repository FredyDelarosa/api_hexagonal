package application

import (
	"rest/src/ide/domain"
	"rest/src/ide/domain/entities"
)

type AddIDEUseCase struct {
	repo domain.IDERepository
}

func NewAddIDEUseCase(repo domain.IDERepository) *AddIDEUseCase {
	return &AddIDEUseCase{repo: repo}
}

func (uc *AddIDEUseCase) Execute(ide *entities.IDE) error {
	return uc.repo.Create(ide)
}
