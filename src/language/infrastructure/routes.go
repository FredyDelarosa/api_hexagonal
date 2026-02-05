package infrastructure

import (
	"rest/src/core/middleware"
	"rest/src/language/application"
	"rest/src/language/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterLanguageRoutes(
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

	// Rutas públicas (GET - solo lectura)
	router.GET("/languages", viewAllController.GetAllLanguages)
	router.GET("/languages/:id", viewByIDController.GetLanguageByID)

	// Rutas protegidas (POST, PUT, DELETE - modificación)
	protectedLanguageGroup := router.Group("/languages")
	protectedLanguageGroup.Use(middleware.AuthMiddleware())
	{
		protectedLanguageGroup.POST("", addController.AddLanguage)
		protectedLanguageGroup.PUT("/:id", updateController.UpdateLanguage)
		protectedLanguageGroup.DELETE("/:id", removeController.RemoveLanguage)
	}
}
