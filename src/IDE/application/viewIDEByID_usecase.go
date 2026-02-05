package application

import (
	"rest/src/IDE/domain"
	"rest/src/IDE/domain/entities"
)

type ViewIDEByIDUseCase struct {
	repo domain.IDERepository
}

func NewViewIDEByIDUseCase(repo domain.IDERepository) *ViewIDEByIDUseCase {
	return &ViewIDEByIDUseCase{
		repo: repo,
	}
}

func (uc *ViewIDEByIDUseCase) Execute(id int) (*entities.IDE, error) {
	return uc.repo.GetByID(id)
}
