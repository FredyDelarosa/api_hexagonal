package controllers

import (
	"net/http"
	"rest/src/ide/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemoveIDEController struct {
	useCase *application.RemoveIDEUseCase
}

func NewRemoveIDEController(useCase *application.RemoveIDEUseCase) *RemoveIDEController {
	return &RemoveIDEController{useCase: useCase}
}

func (ctrl *RemoveIDEController) DeleteIDE(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := ctrl.useCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "IDE eliminado correctamente"})
}
