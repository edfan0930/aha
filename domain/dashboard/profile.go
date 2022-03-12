package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

//Profile
func Profile(c *gin.Context) {

	//get value from session storage
	email := c.Request.Header.Get("email")

	//get user
	user, err := db.First(db.MainSession, c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	//just user defined could reset password
	var reset bool
	if user.Password != "" {
		reset = true
	}

	//template
	c.HTML(http.StatusOK, "profile.html", gin.H{
		"email":      user.Email,
		"name":       user.Name,
		"couldReset": reset,
	})
}
