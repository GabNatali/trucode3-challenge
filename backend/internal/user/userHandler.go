package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	AddUser(c *gin.Context)
}

type userHandler struct {
	userService UserService
}

func NewUserHandler(cases UserService) UserHandler {
	return &userHandler{
		userService: cases,
	}
}

func (u *userHandler) AddUser(c *gin.Context) {

	var addUserDto AddUserDto

	if err := c.ShouldBindJSON(&addUserDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := u.userService.Add(addUserDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"id":       user.Id,
		"message":  "User added successfully",
	})
}
