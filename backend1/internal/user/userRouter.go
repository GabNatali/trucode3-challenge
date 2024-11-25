package user

import (
	"github.com/GabNatali/trucode3-challenge-final/middleware"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.Engine, handlers UserHandler) {
	r.POST("/users", handlers.AddUser)
	r.GET("/users/me", middleware.Authenticate(), handlers.Getme)
	r.PATCH("/users/config", middleware.Authenticate(), handlers.UpdateConfigUser)
}
