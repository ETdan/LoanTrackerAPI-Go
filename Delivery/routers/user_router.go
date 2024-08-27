package routers

import (
	"LoanAPI/LoanTrackerAPI-Go/Delivery/controllers"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(router *gin.RouterGroup, userController controllers.UserController) {
	router.POST("/register", userController.Register)
	router.GET("/verify-email", userController.VerifyEmail)
	router.POST("/login", userController.Login)

	router.POST("/token/refresh", userController.RefreshToken)
	router.POST("/password-reset", userController.PasswordReset)

	router.GET("/users/profile", userController.Profile)

}
