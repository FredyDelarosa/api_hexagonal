package application

import (
	"rest/src/language/domain"
	"rest/src/language/domain/entities"
)

type ViewAllLanguagesUseCase struct {
	repo domain.LanguageRepository
}

func NewViewAllLanguagesUseCase(repo domain.LanguageRepository) *ViewAllLanguagesUseCase {
	return &ViewAllLanguagesUseCase{repo: repo}
}

func (uc *ViewAllLanguagesUseCase) Execute() ([]entities.Language, error) {
	return uc.repo.GetAll()
}
