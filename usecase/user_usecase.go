package usecase

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"
	"errors"
	"fmt"
)

type UserUseCase struct {
	UserRepo domain.UserRepository
}

// CreateUser implements domain.UserUsecase.
func (u UserUseCase) CreateUnVerifiedUser(user domain.UnverifiedUser) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("Invalid user details")
	}

	_, err := u.UserRepo.GetUnverifiedUserByEmail(user.Email)

	if err == nil {
		return errors.New("Email Already exists")
	}

	return u.UserRepo.CreateUnVerifiedUser(user)
}
func (u UserUseCase) VerifyEmail(token string) error {
	email, err := infrastructure.ParseClaims(token)
	if err != nil {
		return err
	}

	if email == "" {
		return errors.New("Invalid token")
	}

	_, err = u.UserRepo.GetUserByEmail(email)

	if err == nil {
		return errors.New("Email Already verified")
	}

	unverified_user, err := u.UserRepo.GetUnverifiedUserByEmail(email)
	if err != nil {
		return errors.New("verification email not found")
	}

	var user domain.User
	user.Email = unverified_user.Email
	user.Password, err = infrastructure.HashPassword(unverified_user.Password)
	if err != nil {
		return err
	}

	return u.UserRepo.CreateUser(user)
}
func (u UserUseCase) Login(email string, password string) (string, string, error) {
	user, err := u.UserRepo.GetUserByEmail(email)
	fmt.Println("User ID: ", user.ID.Hex(), user.Password)
	if err != nil {
		return "", "", err
	}

	if infrastructure.ComparePasswords(user.Password, password) != nil {
		return "", "", errors.New("Invalid email or password")
	}
	accesToken, RefreshToken, err := infrastructure.GenerateAccessTokenAndRefreshToken(user.ID.Hex())
	if err != nil {
		return "", "", err
	}

	return accesToken, RefreshToken, nil
}
func (u UserUseCase) RefreshToken(refreshToken string) (string, error) {
	// Verify the refresh token and extract the user_id
	userID, err := infrastructure.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid refresh token: %w", err)
	}

	// Check if the user exists
	_, err = u.UserRepo.GetUser(userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate a new access token
	accessToken, _, err := infrastructure.GenerateAccessTokenAndRefreshToken(userID)
	if err != nil {
		return "", fmt.Errorf("failed to generate new access token: %w", err)
	}

	return accessToken, nil
}

func (u UserUseCase) ResetPassword(sender_email string, token string, password string) error {
	email, err := infrastructure.ParseClaims(token)
	if err != nil {
		return err
	}

	if email == "" || sender_email == "" {
		return errors.New("Invalid token")
	}

	if password == "" {
		return errors.New("Invalid password")
	}

	user_id, err := infrastructure.VerifyRefreshToken(token)

	if err != nil {
		return err
	}

	user, err := u.UserRepo.GetUser(user_id)
	if err != nil {
		return errors.New("User not found")
	}

	user, err = u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("User not found")
	}

	user.Password, err = infrastructure.HashPassword(password)
	if err != nil {
		return err
	}

	_, err = u.UserRepo.UpdateUser(user.ID.Hex(), user)
	return err
}

func (u UserUseCase) InitiatePasswordReset(email string) error {
	user, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("User not found")
	}

	token, err := infrastructure.GenerateUserJWTToken(user)
	if err != nil {
		return err
	}

	smtpHost := "smtp.gmail.com"              // e.g., "smtp.gmail.com"
	smtpPort := 587                           // e.g., 587 for TLS
	smtpUser := "yordanoslegesse15@gmail.com" // Your SMTP username (often your email address)
	smtpPassword := "bcewmdllhervddxu"        // Your SMTP password

	err = infrastructure.SendPassworkResetEmail(user.Email, token, smtpHost, smtpUser, smtpPassword, smtpPort)
	if err != nil {
		return err
	}

	return nil
}

// GetUser implements domain.UserUsecase.
func (u UserUseCase) GetUser(id string) (domain.User, error) {
	user, err := u.UserRepo.GetUser(id)
	if id == user.ID.Hex() {
		return user, nil
	}
	return user, err
}

// GetUsers implements domain.UserUsecase.
func (u UserUseCase) GetUsers(user_id string) ([]domain.User, error) {
	user, err := u.UserRepo.GetUser(user_id)
	if err != nil {
		return nil, err
	}
	if user.Role != "admin" {
		return nil, errors.New("Unauthorized access to user list")
	}

	users, err := u.UserRepo.GetUsers(user_id)
	return users, err
}

func (u UserUseCase) Profile(user_id string) (domain.User, error) {
	user, err := u.UserRepo.GetUser(user_id)
	if err != nil {
		return domain.User{}, err
	}
	if user.ID.Hex() != user_id {
		return domain.User{}, errors.New("Unauthorized access to user profile")
	}
	return user, err
}

func (u UserUseCase) DeleteUser(id string, id_to_delete string) error {
	user, err := u.UserRepo.GetUser(id)
	if err != nil {
		return err
	}
	if user.Role != "admin" {
		return errors.New("Unauthorized access to delete user")
	}
	_, err = u.UserRepo.GetUser(id_to_delete)
	if err != nil {
		return err
	}

	err = u.UserRepo.DeleteUser(id_to_delete)
	return err
}

// UpdateUser implements domain.UserUsecase.
func (u UserUseCase) UpdateUser(id string, user domain.User) (domain.User, error) {
	updatedUser, err := u.UserRepo.UpdateUser(id, user)
	return updatedUser, err
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return UserUseCase{
		UserRepo: userRepo,
	}
}
