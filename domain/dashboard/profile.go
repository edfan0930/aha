package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	email := c.Request.Header.Get("email")
	//name := c.Request.Header.Get("name")

	user, err := db.First(db.MainSession, c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	var reset bool
	if user.Password != "" {
		reset = true
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"email":      user.Email,
		"name":       user.Name,
		"couldReset": reset,
	})
}
