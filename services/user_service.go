package services

import (
	"github.com/johnwoz123/pharmacy-user-api/domain/users"
	"github.com/johnwoz123/pharmacy-user-api/utils/date"
	"github.com/johnwoz123/pharmacy-user-api/utils/errors"
	"github.com/johnwoz123/pharmacy-user-api/utils/password/crypt"
	"strconv"
)

var UserService userServiceInterface = &userService{}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErrors)
	GetUser(int64) (*users.User, *errors.RestErrors)
	UpdateUser(users.User) (*users.User, *errors.RestErrors)
	DeleteUser(int64) *errors.RestErrors
	FindByStatus(string) (users.Users, *errors.RestErrors)
}

type userService struct {
}

func getUserById(idParam string) (int64, *errors.RestErrors) {
	userId, Uerr := strconv.ParseInt(idParam, 10, 64)
	if Uerr != nil {
		return 0, errors.BadRequestError("user id must bea number")
	}
	return userId, nil
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErrors) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = date.GetNowForDatabase()
	user.Status = users.StatusActive
	user.Password = crypt.FetchMD5(user.Password)
	if err := user.ValidatePassword(); err != nil {
		return nil, err
	}

	if err := user.Persist(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userId int64) (*users.User, *errors.RestErrors) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func (s *userService) UpdateUser(user users.User) (*users.User, *errors.RestErrors) {
	currentUser := &users.User{Id: user.Id}
	if err := currentUser.Get(); err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (s *userService) DeleteUser(userId int64) *errors.RestErrors {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *userService) FindByStatus(status string) (users.Users, *errors.RestErrors) {
	dataAccessObject := &users.User{}
	return dataAccessObject.FindByStatus(status)
}
