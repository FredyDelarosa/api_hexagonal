package controllers

import (
	"rest/src/language/application"
	"rest/src/language/domain/entities"

	"github.com/gin-gonic/gin"
)

type AddLanguageController struct {
	useCase *application.AddLanguageUseCase
}

func NewAddLanguageController(useCase *application.AddLanguageUseCase) *AddLanguageController {
	return &AddLanguageController{useCase: useCase}
}

func (ctrl *AddLanguageController) AddLanguage(c *gin.Context) {
	var language entities.Language
	if err := c.ShouldBindJSON(&language); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.useCase.Execute(&language)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error al guardar el lenguaje"})
		return
	}

	c.JSON(201, gin.H{"message": "Lenguaje agregado exitosamente"})
}
