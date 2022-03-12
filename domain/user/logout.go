package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"

	"github.com/gin-gonic/gin"
)

//Logout delete session storage
func Logout(c *gin.Context) {

	s := storage.NewSession(storage.PassSecure(c.Request))
	s.Logout(c.Writer, c.Request)

	c.Redirect(http.StatusSeeOther, "/login")
}
