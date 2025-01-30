package controllers

import (
	"net/http"
	"strconv"

	"rest/src/language/application"

	"github.com/gin-gonic/gin"
)

type RemoveLanguageController struct {
	useCase *application.RemoveLanguageUseCase
}

func NewRemoveLanguageController(useCase *application.RemoveLanguageUseCase) *RemoveLanguageController {
	return &RemoveLanguageController{useCase: useCase}
}

func (ctrl *RemoveLanguageController) RemoveLanguage(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = ctrl.useCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el lenguaje"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lenguaje eliminado exitosamente"})
}
