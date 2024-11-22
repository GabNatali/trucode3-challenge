package stats

import (
	"github.com/gin-gonic/gin"
)

type StatsHandler interface {
	GetStats(c *gin.Context)
}

type statsHandler struct {
	repo StatsRepository
}

func NewHandler(repo StatsRepository) StatsHandler {
	return &statsHandler{
		repo: repo,
	}
}

func (s *statsHandler) GetStats(c *gin.Context) {

	option := c.Query("option")

	if option != "education" && option != "occupation" {
		c.JSON(400, gin.H{
			"error": "Invalid option",
		})
		return
	}

	stats, err := s.repo.GetStats(option)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stats)
}
