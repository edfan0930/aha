package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

//UserData
func UserData(c *gin.Context) {

	//users statistics
	statistics, err := db.UserStatistics()
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//user list
	userList, err := db.UserList()
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"total_signed_up": statistics.Total,
		"sessions_today":  statistics.SessionsToday,
		"active_average":  statistics.ActiveAVG,
		"list":            userList,
	})
}
