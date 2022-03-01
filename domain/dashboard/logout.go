package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

//
func Logout(c *gin.Context) {

	if err := gothic.Logout(c.Writer, c.Request); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
	}

	c.JSON(http.StatusSeeOther, "/login")
}
