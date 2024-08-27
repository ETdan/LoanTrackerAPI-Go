package usecase

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"errors"
)

type UserUseCase struct {
	UserRepo domain.UserRepository
}

// CreateUser implements domain.UserUsecase.
func (u UserUseCase) CreateUser(user domain.UnverifiedUser) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("Invalid user details")
	}

	_, err := u.UserRepo.GetUserByEmail(user.Email)

	if err == nil {
		return errors.New("Email Already exists")
	}

	return u.UserRepo.CreateUser(user)
}

// DeleteUser implements domain.UserUsecase.
func (u UserUseCase) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements domain.UserUsecase.
func (u UserUseCase) GetUser(id string) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserUsecase.
func (u UserUseCase) GetUsers() ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUsecase.
func (u UserUseCase) UpdateUser(id string, user domain.User) (domain.User, error) {
	panic("unimplemented")
}

func NewUserUseCase(userRepo domain.UserRepository) domain.UserUsecase {
	return UserUseCase{
		UserRepo: userRepo,
	}
}
