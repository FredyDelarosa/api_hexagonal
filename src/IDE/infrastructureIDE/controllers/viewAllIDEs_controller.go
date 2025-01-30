package controllers

import (
	"net/http"
	"rest/src/ide/application"

	"github.com/gin-gonic/gin"
)

type ViewAllIDEsController struct {
	useCase *application.ViewAllIDEsUseCase
}

func NewViewAllIDEsController(useCase *application.ViewAllIDEsUseCase) *ViewAllIDEsController {
	return &ViewAllIDEsController{useCase: useCase}
}

func (ctrl *ViewAllIDEsController) ListAllIDEs(c *gin.Context) {
	ides, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ides)
}
