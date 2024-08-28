package controllers

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"fmt"
	"net/http"

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
	err := u.UserUsecase.CreateUnVerifiedUser(user)
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
	token := c.Query("token")
	if token == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid token",
			Status:  400,
		})
	}
	err := u.UserUsecase.VerifyEmail(token)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Email verified successfully",
		Status:  200,
	})
}

func (u UserController) Login(c *gin.Context) {
	var login domain.Login
	if c.ShouldBind(&login) != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request body",
			Status:  400,
		})
	}

	if login.Email == "" || login.Password == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid email or password",
			Status:  400,
		})
		return
	}

	fmt.Println(login.Email, login.Password, "email and password")
	accessToken, refreshToken, err := u.UserUsecase.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	fmt.Println(accessToken, refreshToken, "access and refresh token")
	c.JSON(200, gin.H{
		"status":        200,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (u UserController) RefreshToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}

	// Bind JSON from the request body
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	newAccessToken, err := u.UserUsecase.RefreshToken(requestBody.RefreshToken)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}
	// Send the new access token as a response
	c.JSON(http.StatusOK, gin.H{
		"refresh_token": newAccessToken,
	})
}

func (u UserController) PasswordReset(c *gin.Context) {
	// Extract token and email from the request body
	var request struct {
		Email       string `json:"email" binding:"required,email"`
		NewPassword string `json:"new_password"`
	}

	token := c.Query("token")

	// Bind the JSON request to the request struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// If a token is provided, attempt to reset the password
	if token != "" {
		err := u.UserUsecase.ResetPassword(request.Email, token, request.NewPassword)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Respond with success if the password was successfully reset
		c.JSON(http.StatusOK, gin.H{"message": "Password successfully reset"})
		return
	}

	// If no token is provided, initiate the password reset process by sending a reset link to the email
	err := u.UserUsecase.InitiatePasswordReset(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate password reset"})
		return
	} else {
		// Respond with success if the reset link was successfully sent
		c.JSON(http.StatusOK, gin.H{"message": "Password reset instructions sent to email"})
	}
}

func (u UserController) Profile(c *gin.Context) {
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid user ID",
			Status:  400,
		})
		return
	}
	user, err := u.UserUsecase.Profile(user_id)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 200,
		"user":   user,
	})
}

func (u UserController) DeleteUser(c *gin.Context) {
	user_id := c.GetString("user_id")
	delete_user_id := c.Param("id")
	if delete_user_id == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid user ID",
			Status:  400,
		})
		return
	}
	if user_id == "" {
		c.JSON(400, domain.ErrorResponse{
			Message: "unauthorized access to delete user",
			Status:  400,
		})
		return
	}
	err := u.UserUsecase.DeleteUser(user_id, delete_user_id)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "User deleted successfully",
		Status:  200,
	})
}

func (u UserController) GetAllUsers(c *gin.Context) {
	user_id := c.GetString("user_id")
	users, err := u.UserUsecase.GetUsers(user_id)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Users retrieved successfully",
		Status:  200,
		Data:    users,
	})
}

// loan controller

func (u UserController) ApplyLoan(c *gin.Context) {
	var loan domain.Loan
	if c.ShouldBind(&loan) != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request body",
			Status:  400,
		})
		return
	}
	err := u.UserUsecase.ApplyLoan(loan)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}

	response := fmt.Sprintf("Loan application was successful")
	c.JSON(201, domain.SuccessResponse{
		Message: response,
		Status:  201,
	})
}

func (u UserController) GetLoans(c *gin.Context) {
	user_id := c.GetString("user_id")
	loans, err := u.UserUsecase.GetLoans(user_id)
	if err != nil {
		c.JSON(500, domain.ErrorResponse{
			Message: err.Error(),
			Status:  500,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Loans retrieved successfully",
		Status:  200,
		Data:    loans,
	})
}
