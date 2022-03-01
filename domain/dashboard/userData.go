package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

//UserStatistics
func UserData(c *gin.Context) {

	statistics, err := db.UserStatistics()
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"total_signed_up": statistics.Total,
		"session_today":   statistics.SessionToday,
		"activ_average":   statistics.ActiveAVG,
	})
}
