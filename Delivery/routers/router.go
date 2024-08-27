package routers

import (
	"LoanAPI/LoanTrackerAPI-Go/Delivery/controllers"
	"LoanAPI/LoanTrackerAPI-Go/repository"
	"LoanAPI/LoanTrackerAPI-Go/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RouterSetUp(router *gin.Engine, db *mongo.Database) {
	usersRouter := router.Group("users")
	userRepository := repository.NewUserRepository(*db.Collection("user"), *db.Collection("unverified_user"))

	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controllers.NewUsersController(userUseCase)
	NewUserRouter(usersRouter, userController)

	adminRouter := router.Group("admin")
	adminRepository := repository.NewUserRepository(*db.Collection("user"), *db.Collection("unverified_user"))
	adminUseCase := usecase.NewUserUseCase(adminRepository)
	adminController := controllers.NewUsersController(adminUseCase)
	NewAdminRouter(adminRouter, adminController)

}
