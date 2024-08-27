package controllers

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUsersController(userUsecase domain.UserUsecase) UserController {
	return UserController{
		UserUsecase: userUsecase,
	}
}

func (u UserController) Register(c *gin.Context) {
	var user domain.UnverifiedUser

	if c.ShouldBind(&user) != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request body",
			Status:  400,
		})
		return
	}
	err := u.UserUsecase.CreateUser(user)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	response := fmt.Sprintf("Verification email was sent to %s", user.Email)
	c.JSON(201, domain.SuccessResponse{
		Message: response,
		Status:  201,
	})
}

func (u UserController) VerifyEmail(c *gin.Context) {

}

func (u UserController) Login(c *gin.Context) {
	panic("unimplemented")
}

func (u UserController) RefreshToken(c *gin.Context) {
	panic("unimplemented")
}

func (u UserController) PasswordReset(c *gin.Context) {
	panic("unimplemented")
}

func (u UserController) Profile(c *gin.Context) {
	panic("unimplemented")
}
