package application

import (
	"rest/src/language/domain"
	"rest/src/language/domain/entities"
)

type ViewLanguageByIDUseCase struct {
	repo domain.LanguageRepository
}

func NewViewLanguageByIDUseCase(repo domain.LanguageRepository) *ViewLanguageByIDUseCase {
	return &ViewLanguageByIDUseCase{repo: repo}
}

func (uc *ViewLanguageByIDUseCase) Execute(id int) (*entities.Language, error) {
	return uc.repo.GetByID(id)
}
