package controllers

import (
	"net/http"
	"rest/src/auth/application"

	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	useCase *application.LogoutUseCase
}

func NewLogoutController(useCase *application.LogoutUseCase) *LogoutController {
	return &LogoutController{useCase: useCase}
}

type LogoutRequest struct {
	UserID int `json:"user_id" binding:"required"`
}

func (ctrl *LogoutController) Logout(c *gin.Context) {
	var req LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.useCase.Execute(req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Sesión cerrada exitosamente",
	})
}
