package routers

import (
	"LoanAPI/LoanTrackerAPI-Go/Delivery/controllers"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewAdminRouter(router *gin.RouterGroup, userController controllers.UserController) {

	router.DELETE("/delete/:id", infrastructure.Auth_middleware(), userController.DeleteUser)
	router.GET("/users", infrastructure.Auth_middleware(), userController.GetAllUsers)
}
