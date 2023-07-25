package requests

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserCreateRequest struct {
	Name     string `json:"name,omitempty" validate:"required"`
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

func (ucr UserCreateRequest) EncryptPassword() string {
	password, err := bcrypt.GenerateFromPassword([]byte(ucr.Password), 14)
	if err != nil {
		log.Fatal("Error in encrypting the password")
	}
	return string(password)
}
