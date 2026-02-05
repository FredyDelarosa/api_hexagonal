package infrastructure

import (
	"rest/src/auth/application"
	"rest/src/auth/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(
	router *gin.Engine,
	registerUseCase *application.RegisterUseCase,
	loginUseCase *application.LoginUseCase,
	logoutUseCase *application.LogoutUseCase,
) {
	registerController := controllers.NewRegisterController(registerUseCase)
	loginController := controllers.NewLoginController(loginUseCase)
	logoutController := controllers.NewLogoutController(logoutUseCase)

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", registerController.Register)
		authGroup.POST("/login", loginController.Login)
		authGroup.POST("/logout", logoutController.Logout)
	}
}
