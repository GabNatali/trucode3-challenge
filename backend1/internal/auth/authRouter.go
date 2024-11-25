package auth

import "github.com/gin-gonic/gin"

func AddAuthRouter(r *gin.Engine, handler AuthHandler) {

	authGroup := r.Group("/auth")
	authGroup.POST("/sign-in", handler.Login)
}
