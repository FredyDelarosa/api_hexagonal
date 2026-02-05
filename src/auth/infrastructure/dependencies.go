package infrastructure

import (
	"rest/src/auth/application"
	"rest/src/core"
)

type Dependencies struct {
	RegisterUseCase *application.RegisterUseCase
	LoginUseCase    *application.LoginUseCase
	LogoutUseCase   *application.LogoutUseCase
}

func NewDependencies() (*Dependencies, error) {
	db, err := core.InitDB()
	if err != nil {
		return nil, err
	}

	userRepo := NewMySQLUserRepository(db)

	return &Dependencies{
		RegisterUseCase: application.NewRegisterUseCase(userRepo),
		LoginUseCase:    application.NewLoginUseCase(userRepo),
		LogoutUseCase:   application.NewLogoutUseCase(),
	}, nil
}
