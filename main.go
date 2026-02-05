package main

import (
	"log"
	"rest/src/IDE/infrastructureIDE"
	authInfra "rest/src/auth/infrastructure"
	"rest/src/language/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	langDeps, errLang := infrastructure.NewDependencies()
	ideDeps, errIDE := infrastructureIDE.NewDependencies()
	authDeps, errAuth := authInfra.NewDependencies()

	if errLang != nil {
		log.Fatal("Error al inicializar dependencias de Language:", errLang)
	}
	if errIDE != nil {
		log.Fatal("Error al inicializar dependencias de IDE:", errIDE)
	}
	if errAuth != nil {
		log.Fatal("Error al inicializar dependencias de Auth:", errAuth)
	}

	r := gin.Default()

	// Registrar rutas
	infrastructure.RegisterLanguageRoutes(r, langDeps.AddLanguageUseCase, langDeps.ViewAllLanguagesUseCase, langDeps.ViewLanguageByIDUseCase, langDeps.RemoveLanguageUseCase, langDeps.UpdateLanguageUseCase)

	infrastructureIDE.RegisterIDERoutes(r, ideDeps.AddIDEUseCase, ideDeps.ViewAllIDEsUseCase, ideDeps.ViewIDEByIDUseCase, ideDeps.RemoveIDEUseCase, ideDeps.UpdateIDEUseCase)

	authInfra.RegisterAuthRoutes(r, authDeps.RegisterUseCase, authDeps.LoginUseCase, authDeps.LogoutUseCase)

	r.Run(":8080")
}
