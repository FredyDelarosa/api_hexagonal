package infrastructure

import (
	"rest/src/language/application"
	"rest/src/language/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	addLanguageUseCase *application.AddLanguageUseCase,
	viewAllLanguagesUseCase *application.ViewAllLanguagesUseCase,
	viewLanguageByIDUseCase *application.ViewLanguageByIDUseCase,
	removeLanguageUseCase *application.RemoveLanguageUseCase,
	updateLanguageUseCase *application.UpdateLanguageUseCase,
) {
	addController := controllers.NewAddLanguageController(addLanguageUseCase)
	viewAllController := controllers.NewViewAllLanguagesController(viewAllLanguagesUseCase)
	viewByIDController := controllers.NewViewLanguageByIDController(viewLanguageByIDUseCase)
	removeController := controllers.NewRemoveLanguageController(removeLanguageUseCase)
	updateController := controllers.NewUpdateLanguageController(updateLanguageUseCase)

	router.POST("/languages", addController.AddLanguage)
	router.GET("/languages", viewAllController.GetAllLanguages)
	router.GET("/languages/:id", viewByIDController.GetLanguageByID)
	router.DELETE("/languages/:id", removeController.RemoveLanguage)
	router.PUT("/languages/:id", updateController.UpdateLanguage)
}
