package controllers

import (
	"net/http"
	"rest/src/ide/application"
	"rest/src/ide/domain/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateIDEController struct {
	useCase *application.UpdateIDEUseCase
}

func NewUpdateIDEController(useCase *application.UpdateIDEUseCase) *UpdateIDEController {
	return &UpdateIDEController{useCase: useCase}
}

func (ctrl *UpdateIDEController) UpdateIDE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var ide entities.IDE
	if err := c.ShouldBindJSON(&ide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ide.ID = id
	if err := ctrl.useCase.Execute(&ide); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "IDE actualizado correctamente"})
}
