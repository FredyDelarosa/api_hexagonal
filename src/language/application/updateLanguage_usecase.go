package application

import (
	"rest/src/language/domain"
	"rest/src/language/domain/entities"
)

type UpdateLanguageUseCase struct {
	repo domain.LanguageRepository
}

func NewUpdateLanguageUseCase(repo domain.LanguageRepository) *UpdateLanguageUseCase {
	return &UpdateLanguageUseCase{repo: repo}
}

func (uc *UpdateLanguageUseCase) Execute(language *entities.Language) error {
	return uc.repo.Update(language)
}
