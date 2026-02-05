package infrastructureIDE

import (
"rest/src/IDE/application"
"rest/src/IDE/infrastructureIDE/controllers"
"rest/src/core/middleware"

"github.com/gin-gonic/gin"
)

func RegisterIDERoutes(
router *gin.Engine,
addIDEUseCase *application.AddIDEUseCase,
viewAllIDEsUseCase *application.ViewAllIDEsUseCase,
viewIDEByIDUseCase *application.ViewIDEByIDUseCase,
removeIDEUseCase *application.RemoveIDEUseCase,
updateIDEUseCase *application.UpdateIDEUseCase,
) {
	addController := controllers.NewAddIDEController(addIDEUseCase)
	viewAllController := controllers.NewViewAllIDEsController(viewAllIDEsUseCase)
	viewByIDController := controllers.NewViewIDEByIDController(viewIDEByIDUseCase)
	removeController := controllers.NewRemoveIDEController(removeIDEUseCase)
	updateController := controllers.NewUpdateIDEController(updateIDEUseCase)

	// Rutas públicas (GET - solo lectura)
	router.GET("/ides", viewAllController.ListAllIDEs)
	router.GET("/ides/:id", viewByIDController.GetIDEByID)

	// Rutas protegidas (POST, PUT, DELETE - modificación)
	protectedIdeGroup := router.Group("/ides")
	protectedIdeGroup.Use(middleware.AuthMiddleware())
	{
		protectedIdeGroup.POST("", addController.CreateIDE)
		protectedIdeGroup.PUT("/:id", updateController.UpdateIDE)
		protectedIdeGroup.DELETE("/:id", removeController.DeleteIDE)
	}
}
