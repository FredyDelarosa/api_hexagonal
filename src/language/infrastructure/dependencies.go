package infrastructure

import (
	"rest/src/core"
	"rest/src/language/application"
)

type Dependencies struct {
	AddLanguageUseCase      *application.AddLanguageUseCase
	ViewAllLanguagesUseCase *application.ViewAllLanguagesUseCase
	ViewLanguageByIDUseCase *application.ViewLanguageByIDUseCase
	RemoveLanguageUseCase   *application.RemoveLanguageUseCase
	UpdateLanguageUseCase   *application.UpdateLanguageUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	mysqlRepo := NewMySQLRepository(db)

	return &Dependencies{
		AddLanguageUseCase:      application.NewAddLanguageUseCase(mysqlRepo),
		ViewAllLanguagesUseCase: application.NewViewAllLanguagesUseCase(mysqlRepo),
		ViewLanguageByIDUseCase: application.NewViewLanguageByIDUseCase(mysqlRepo),
		RemoveLanguageUseCase:   application.NewRemoveLanguageUseCase(mysqlRepo),
		UpdateLanguageUseCase:   application.NewUpdateLanguageUseCase(mysqlRepo),
	}, nil
}
