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
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user UnverifiedUser) error
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
}

type UserRepository interface {
	GetUsers() ([]User, error)
	GetUser(id string) (User, error)
	CreateUser(user UnverifiedUser) error
	UpdateUser(id string, user User) (User, error)
	DeleteUser(id string) error
	GetUserByEmail(email string) (User, error)
}
