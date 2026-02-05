package controllers

import (
	"net/http"
	"rest/src/auth/application"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	useCase *application.LoginUseCase
}

func NewLoginController(useCase *application.LoginUseCase) *LoginController {
	return &LoginController{useCase: useCase}
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ctrl *LoginController) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := ctrl.useCase.Execute(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login exitoso",
		"token":   loginResponse.Token,
		"user": gin.H{
			"id":         loginResponse.User.ID,
			"email":      loginResponse.User.Email,
			"first_name": loginResponse.User.FirstName,
			"last_name":  loginResponse.User.LastName,
		},
	})
}
