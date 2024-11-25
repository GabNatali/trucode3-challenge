package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	AddUser(c *gin.Context)
	Getme(c *gin.Context)
	UpdateConfigUser(c *gin.Context)
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
		"data":    user,
		"message": "User added successfully",
		"status":  http.StatusCreated,
	})
}

func (u *userHandler) Getme(c *gin.Context) {

	token := c.GetHeader("Authorization")

	user, err := u.userService.GetMe(token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)

}

type UpdateBody struct {
	UserId        uint   `json:"user_id"`
	Education     string `json:"education"`
	Income        string `json:"income"`
	MaritalStatus string `json:"marital_status"`
	MaxAge        int    `json:"max_age"`
	MinAge        int    `json:"min_age"`
	Occupation    string `json:"occupation"`
	OrderBy       string `json:"order_by"`
	OrderDir      string `json:"order_dir"`
}

func (u *userHandler) UpdateConfigUser(c *gin.Context) {

	var updateBody UpdateBody

	if err := c.ShouldBindJSON(&updateBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("updateBody")
	fmt.Println(updateBody)
	user, err := u.userService.UpdateConfig(updateBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
