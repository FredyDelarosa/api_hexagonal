package application

import "rest/src/language/domain"

type RemoveLanguageUseCase struct {
	repo domain.LanguageRepository
}

func NewRemoveLanguageUseCase(repo domain.LanguageRepository) *RemoveLanguageUseCase {
	return &RemoveLanguageUseCase{repo: repo}
}

func (uc *RemoveLanguageUseCase) Execute(id int) error {
	return uc.repo.Delete(id)
}
