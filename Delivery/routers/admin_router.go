package routers

import (
	"LoanAPI/LoanTrackerAPI-Go/Delivery/controllers"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"

	"github.com/gin-gonic/gin"
)

func NewAdminRouter(router *gin.RouterGroup, userController controllers.UserController) {

	router.DELETE("/delete/:id", infrastructure.Auth_middleware(), userController.DeleteUser)
	router.GET("/users", infrastructure.Auth_middleware(), userController.GetAllUsers)
	router.GET("loans", infrastructure.Auth_middleware(), userController.GetLoans)

	router.PUT("/loans/:id/:status", infrastructure.Auth_middleware(), userController.UpdateLoan)
	router.DELETE("/loans/:id", infrastructure.Auth_middleware(), userController.DeleteLoan)

	router.GET("logs", infrastructure.Auth_middleware(), controllers.Log())

}
