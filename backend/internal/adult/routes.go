package adult

import (
	"github.com/GabNatali/trucode3-challenge-final/middleware"
	"github.com/gin-gonic/gin"
)

func AddRoutesAdult(router *gin.Engine, handlers AdultHandler) {

	router.Use(middleware.Authenticate())
	router.GET("/adults", handlers.Find)
	router.GET("/adults/export", handlers.ExportCSV)
}
