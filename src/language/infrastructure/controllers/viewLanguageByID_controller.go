package controllers

import (
	"net/http"
	"strconv"

	"rest/src/language/application"

	"github.com/gin-gonic/gin"
)

type ViewLanguageByIDController struct {
	useCase *application.ViewLanguageByIDUseCase
}

func NewViewLanguageByIDController(useCase *application.ViewLanguageByIDUseCase) *ViewLanguageByIDController {
	return &ViewLanguageByIDController{useCase: useCase}
}

func (ctrl *ViewLanguageByIDController) GetLanguageByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	language, err := ctrl.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el lenguaje"})
		return
	}

	if language == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Lenguaje no encontrado"})
		return
	}

	c.JSON(http.StatusOK, language)
}
