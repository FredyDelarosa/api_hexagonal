package infrastructureIDE

import (
	"rest/src/core"
	"rest/src/ide/application"
)

type Dependencies struct {
	AddIDEUseCase      *application.AddIDEUseCase
	ViewAllIDEsUseCase *application.ViewAllIDEsUseCase
	ViewIDEByIDUseCase *application.ViewIDEByIDUseCase
	RemoveIDEUseCase   *application.RemoveIDEUseCase
	UpdateIDEUseCase   *application.UpdateIDEUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	mysqlRepo := NewMySQLRepository(db)

	return &Dependencies{
		AddIDEUseCase:      application.NewAddIDEUseCase(mysqlRepo),
		ViewAllIDEsUseCase: application.NewViewAllIDEsUseCase(mysqlRepo),
		ViewIDEByIDUseCase: application.NewViewIDEByIDUseCase(mysqlRepo),
		RemoveIDEUseCase:   application.NewRemoveIDEUseCase(mysqlRepo),
		UpdateIDEUseCase:   application.NewUpdateIDEUseCase(mysqlRepo),
	}, nil
}
