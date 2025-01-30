package main

import (
	"log"
	"rest/src/ide/infrastructureIDE"
	"rest/src/language/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	langDeps, errLang := infrastructure.NewDependencies()

	ideDeps, errIDE := infrastructureIDE.NewDependencies()

	if errLang != nil {
		log.Fatal("Error al inicializar dependencias de Language:", errLang)
	}
	if errIDE != nil {
		log.Fatal("Error al inicializar dependencias de IDE:", errIDE)
	}

	r := gin.Default()

	infrastructure.RegisterRoutes(r, langDeps.AddLanguageUseCase, langDeps.ViewAllLanguagesUseCase, langDeps.ViewLanguageByIDUseCase, langDeps.RemoveLanguageUseCase, langDeps.UpdateLanguageUseCase)

	infrastructureIDE.RegisterIDERoutes(r, ideDeps.AddIDEUseCase, ideDeps.ViewAllIDEsUseCase, ideDeps.ViewIDEByIDUseCase, ideDeps.RemoveIDEUseCase, ideDeps.UpdateIDEUseCase)

	r.Run(":8080")
}
