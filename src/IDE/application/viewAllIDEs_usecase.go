package application

import (
	"rest/src/ide/domain"
	"rest/src/ide/domain/entities"
)

type ViewAllIDEsUseCase struct {
	repo domain.IDERepository
}

func NewViewAllIDEsUseCase(repo domain.IDERepository) *ViewAllIDEsUseCase {
	return &ViewAllIDEsUseCase{
		repo: repo,
	}
}

func (uc *ViewAllIDEsUseCase) Execute() ([]entities.IDE, error) {
	return uc.repo.GetAll()
}
