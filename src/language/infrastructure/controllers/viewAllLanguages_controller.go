package controllers

import (
	"net/http"
	"rest/src/language/application"

	"github.com/gin-gonic/gin"
)

type ViewAllLanguagesController struct {
	useCase *application.ViewAllLanguagesUseCase
}

func NewViewAllLanguagesController(useCase *application.ViewAllLanguagesUseCase) *ViewAllLanguagesController {
	return &ViewAllLanguagesController{useCase: useCase}
}

func (ctrl *ViewAllLanguagesController) GetAllLanguages(c *gin.Context) {
	languages, err := ctrl.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los lenguajes"})
		return
	}
	c.JSON(http.StatusOK, languages)
}
