package adult

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdultHandler struct {
	service AdultService
}

func NewUserHandler(service AdultService) *AdultHandler {
	return &AdultHandler{
		service: service,
	}
}

type QueryParams struct {
	Education     string `form:"education"`
	MaritalStatus string `form:"marital_status"`
	Occupation    string `form:"occupation"`
	MinAge        int    `form:"min_age"`
	MaxAge        int    `form:"max_age"`
	IncomeLevel   string `form:"income_level"`
	OrderBy       string `form:"order_by" binding:"required"`
	OrderDir      string `form:"order_dir"`
	Page          int    `form:"page"`
	PageSize      int    `form:"page_size"`
}

func (a *AdultHandler) Find(c *gin.Context) {

	var params QueryParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PageSize == 0 {
		params.PageSize = 10
	}

	adults, err := a.service.Find(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(adults)
	c.JSON(http.StatusOK, gin.H{"data": adults})
}

func (a *AdultHandler) ExportCSV(c *gin.Context) {

	var params QueryParams

	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	csvData, err := a.service.ExportCSV(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename=adults.csv")
	c.Writer.Write(csvData)

}
