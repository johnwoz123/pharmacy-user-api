package users

import (
	"github.com/johnwoz123/pharmacy-user-api/utils/errors"
	"strings"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type Users []User

func (user *User) Validate() *errors.RestErrors {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	if user.Email == "" {
		return errors.BadRequestError("invalid email!")
	}

	if user.FirstName == "" {
		return errors.BadRequestError("invalid first_name!")
	}

	if user.LastName == "" {
		return errors.BadRequestError("invalid last_name!")
	}

	return nil
}

func (user *User) ValidatePassword() *errors.RestErrors {
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.BadRequestError("password cannot be blank!")
	}

	if len(user.Password) < 6 {
		return errors.BadRequestError("password must be at least 6 characters!")
	}

	return nil
}
