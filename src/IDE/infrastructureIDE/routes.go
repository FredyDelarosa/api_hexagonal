package infrastructureIDE

import (
	"rest/src/ide/application"
	"rest/src/ide/infrastructureIDE/controllers"

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

	router.POST("/ides", addController.CreateIDE)
	router.GET("/ides", viewAllController.ListAllIDEs)
	router.GET("/ides/:id", viewByIDController.GetIDEByID)
	router.DELETE("/ides/:id", removeController.DeleteIDE)
	router.PUT("/ides/:id", updateController.UpdateIDE)

}
