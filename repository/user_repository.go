package repository

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	UserCollection           mongo.Collection
	UnverifiedUserCollection mongo.Collection
}

// CreateUser implements domain.UserRepository.
func (u UserRepository) CreateUser(user domain.UnverifiedUser) error {

	smtpHost := "smtp.gmail.com"              // e.g., "smtp.gmail.com"
	smtpPort := 587                           // e.g., 587 for TLS
	smtpUser := "yordanoslegesse15@gmail.com" // Your SMTP username (often your email address)
	smtpPassword := "bcewmdllhervddxu"        // Your SMTP password

	// Generate OTP
	token := infrastructure.GenerateVerificationToken()

	// Send OTP via email
	err := infrastructure.SendVerificationEmail(user.Email, token, smtpHost, smtpUser, smtpPassword, smtpPort)
	if err != nil {
		return err
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	_, err = u.UnverifiedUserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.UserRepository.
func (u UserRepository) DeleteUser(id string) error {
	panic("unimplemented")
}

// GetUser implements domain.UserRepository.
func (u UserRepository) GetUser(id string) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserRepository.
func (u UserRepository) GetUsers() ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserRepository.
func (u UserRepository) UpdateUser(id string, user domain.User) (domain.User, error) {
	panic("unimplemented")
}

func (u UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(user)
	return user, err
}

func NewUserRepository(userCollection mongo.Collection, UnverifiedUserCollection mongo.Collection) domain.UserRepository {
	return UserRepository{
		UserCollection:           userCollection,
		UnverifiedUserCollection: UnverifiedUserCollection,
	}
}
