package stats

import (
	"github.com/GabNatali/trucode3-challenge-final/middleware"
	"github.com/gin-gonic/gin"
)

func AddStatsRouter(r *gin.Engine, handlers StatsHandler) {

	r.Use(middleware.Authenticate())
	r.GET("/stats", handlers.GetStats)
}
