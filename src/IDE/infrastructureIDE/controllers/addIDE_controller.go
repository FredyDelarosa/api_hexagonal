package controllers

import (
	"net/http"
	"rest/src/IDE/application"
	"rest/src/IDE/domain/entities"

	"github.com/gin-gonic/gin"
)

type AddIDEController struct {
	useCase *application.AddIDEUseCase
}

func NewAddIDEController(useCase *application.AddIDEUseCase) *AddIDEController {
	return &AddIDEController{useCase: useCase}
}

func (ctrl *AddIDEController) CreateIDE(c *gin.Context) {
	var ide entities.IDE
	if err := c.ShouldBindJSON(&ide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.useCase.Execute(&ide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "IDE agregado exitosamente"})
}
