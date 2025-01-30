package controllers

import (
	"net/http"
	"strconv"

	"rest/src/language/application"
	"rest/src/language/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateLanguageController struct {
	useCase *application.UpdateLanguageUseCase
}

func NewUpdateLanguageController(useCase *application.UpdateLanguageUseCase) *UpdateLanguageController {
	return &UpdateLanguageController{useCase: useCase}
}

func (ctrl *UpdateLanguageController) UpdateLanguage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var language entities.Language
	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	language.ID = id

	err = ctrl.useCase.Execute(&language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el lenguaje"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lenguaje actualizado exitosamente"})
}
