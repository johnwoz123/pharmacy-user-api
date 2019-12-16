package user_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/johnwoz123/pharmacy-user-api/domain/users"
	"github.com/johnwoz123/pharmacy-user-api/services"
	"github.com/johnwoz123/pharmacy-user-api/utils/errors"
	"net/http"
	"strconv"
)

/**
 * Entry into the system
 */
func getUserById(idParam string) (int64, *errors.RestErrors) {
	userId, Uerr := strconv.ParseInt(idParam, 10, 64)
	if Uerr != nil {
		return 0, errors.BadRequestError("user id must bea number")
	}
	return userId, nil
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError("bad request")
		c.JSON(restError.Status, restError)
		return
	}
	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUsers(c *gin.Context) {
	userId, Uerr := getUserById(c.Param("user_id"))
	if Uerr != nil {
		err := errors.BadRequestError("wrong user id")
		c.JSON(err.Status, err)
		return
	}

	result, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result.Encode(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	existingUserId, Uerr := getUserById(c.Param("user_id"))
	if Uerr != nil {
		err := errors.BadRequestError("wrong user id")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.BadRequestError("bad request")
		c.JSON(restError.Status, restError)
		return
	}
	user.Id = existingUserId

	result, updateErr := services.UserService.UpdateUser(user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, idErr := getUserById(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	listOfUsers, err := services.UserService.FindByStatus(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, listOfUsers.Encode(c.GetHeader("X-Public") == "true"))
}
