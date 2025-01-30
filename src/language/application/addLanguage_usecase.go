package application

import (
	"rest/src/language/domain"
	"rest/src/language/domain/entities"
)

type AddLanguageUseCase struct {
	repo domain.LanguageRepository
}

func NewAddLanguageUseCase(repo domain.LanguageRepository) *AddLanguageUseCase {
	return &AddLanguageUseCase{repo: repo}
}

func (uc *AddLanguageUseCase) Execute(language *entities.Language) error {
	return uc.repo.Create(language)
}
