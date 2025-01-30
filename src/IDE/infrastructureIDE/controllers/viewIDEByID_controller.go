package controllers

import (
	"net/http"
	"rest/src/ide/application"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ViewIDEByIDController struct {
	useCase *application.ViewIDEByIDUseCase
}

func NewViewIDEByIDController(useCase *application.ViewIDEByIDUseCase) *ViewIDEByIDController {
	return &ViewIDEByIDController{useCase: useCase}
}

func (ctrl *ViewIDEByIDController) GetIDEByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	ide, err := ctrl.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "IDE no encontrado"})
		return
	}

	c.JSON(http.StatusOK, ide)
}
