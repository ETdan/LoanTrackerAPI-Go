package infrastructure

import "github.com/google/uuid"

func GenerateVerificationToken() string {
	token := uuid.New().String()
	return token
}
