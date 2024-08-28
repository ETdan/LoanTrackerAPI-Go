package repository

import (
	domain "LoanAPI/LoanTrackerAPI-Go/Domain"
	"LoanAPI/LoanTrackerAPI-Go/infrastructure"
	"context"
	"errors"
	"fmt"
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

func (u UserRepository) CreateUser(user domain.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("Invalid user details")
	}

	_, err := u.GetUserByEmail(user.Email)

	if err == nil {
		return errors.New("Email Already exists")
	}
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.Role = "user"

	_, err = u.UserCollection.InsertOne(context.TODO(), user)
	_, err = u.UnverifiedUserCollection.DeleteOne(context.TODO(), bson.M{"email": user.Email})

	if err != nil {
		response := fmt.Sprintf("Error creating user: %s", err.Error())
		return errors.New(response)
	}
	return nil
}

func (u UserRepository) CreateUnVerifiedUser(user domain.UnverifiedUser) error {

	smtpHost := "smtp.gmail.com"              // e.g., "smtp.gmail.com"
	smtpPort := 587                           // e.g., 587 for TLS
	smtpUser := "yordanoslegesse15@gmail.com" // Your SMTP username (often your email address)
	smtpPassword := "bcewmdllhervddxu"        // Your SMTP password

	// Generate OTP
	user.CreatedAt = time.Now()
	token, err := infrastructure.GenerateJWTToken(user)
	if err != nil {
		return err
	}
	// Send OTP via email
	err = infrastructure.SendVerificationEmail(user.Email, token, smtpHost, smtpUser, smtpPassword, smtpPort)
	if err != nil {
		return err
	}
	user.ID = primitive.NewObjectID()
	user.Token = token

	_, err = u.UnverifiedUserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser implements domain.UserRepository.
func (u UserRepository) DeleteUser(id string) error {
	user_object_id, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return errors.New("Invalid user ID")
	}

	_, err = u.UserCollection.DeleteOne(context.TODO(), bson.M{"_id": user_object_id})
	return err
}

// GetUser implements domain.UserRepository.
func (u UserRepository) GetUser(id string) (domain.User, error) {
	user_object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, errors.New("Invalid user ID")
	}

	var user domain.User

	err = u.UserCollection.FindOne(context.TODO(), bson.M{"_id": user_object_id}).Decode(&user)
	return user, err
}

// GetUsers implements domain.UserRepository.
func (u UserRepository) GetUsers(user_id string) ([]domain.User, error) {
	var users []domain.User
	cursor, err := u.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (u UserRepository) UpdateUser(id string, user domain.User) (domain.User, error) {
	user_object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, errors.New("Invalid user ID")
	}

	_, err = u.UserCollection.UpdateOne(context.TODO(), bson.M{"_id": user_object_id}, bson.M{"$set": user})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
func (u UserRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.UserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}
func (u UserRepository) GetUnverifiedUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := u.UnverifiedUserCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}
func NewUserRepository(userCollection mongo.Collection, UnverifiedUserCollection mongo.Collection) domain.UserRepository {
	return UserRepository{
		UserCollection:           userCollection,
		UnverifiedUserCollection: UnverifiedUserCollection,
	}
}
