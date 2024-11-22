package user

import "github.com/gin-gonic/gin"

func AddUserRouter(r *gin.Engine, handlers UserHandler) {
	r.POST("/users", handlers.AddUser)
}
