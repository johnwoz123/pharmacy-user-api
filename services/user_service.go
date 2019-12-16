package services

import (
	"github.com/johnwoz123/pharmacy-user-api/domain/users"
	"github.com/johnwoz123/pharmacy-user-api/utils/date"
	"github.com/johnwoz123/pharmacy-user-api/utils/errors"
	"github.com/johnwoz123/pharmacy-user-api/utils/password/crypt"
)

func CreateUser(user users.User) (*users.User, *errors.RestErrors) {
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

func GetUser(userId int64) (*users.User, *errors.RestErrors) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}

func UpdateUser(user users.User) (*users.User, *errors.RestErrors) {
	currentUser, err := GetUser(user.Id)
	if err != nil {
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

func DeleteUser(userId int64) *errors.RestErrors {
	user := &users.User{Id: userId}
	return user.Delete()
}

func FindByStatus(status string) (users.Users, *errors.RestErrors) {
	dataAccessObject := &users.User{}
	return dataAccessObject.FindByStatus(status)
}
