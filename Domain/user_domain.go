package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserName  string             `json:"user_name" bson:"user_name"`
	Email     string             `json:"email" bson:"email"`
	Role      string             `json:"role" bson:"role"`
	FirstName string             `json:"first_name" bson:"first_name"`
	LastName  string             `json:"last_name" bson:"last_name"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type UnverifiedUser struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	Token     string             `json:"token" bson:"token"`
}

type UserUsecase interface {
	CreateUnVerifiedUser(unverified_user UnverifiedUser) error

	Login(email string, password string) (string, string, error)

	RefreshToken(refresh_token string) (string, error)

	InitiatePasswordReset(email string) error
	ResetPassword(sender_email string, token string, password string) error

	VerifyEmail(token string) error
	GetUsers(user_id string) ([]User, error)
	GetUser(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string, id_to_delete string) error
	Profile(user_id string) (User, error)

	// CreateLoan(loan Loan) error
	ApplyLoan(loan Loan) error
}

type UserRepository interface {
	CreateUser(user User) error
	CreateUnVerifiedUser(unverified_user UnverifiedUser) error

	GetUsers(user_id string) ([]User, error)
	GetUser(id string) (User, error)
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
	GetUserByEmail(email string) (User, error)
	GetUnverifiedUserByEmail(email string) (User, error)
}
