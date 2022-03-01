package callback

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

//Oauth2
func Oauth2(c *gin.Context) {

	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	storage.NewSession(storage.PassSecure(c.Request)).Login(c.Writer, c.Request, gothUser.Email, gothUser.Name, true)
	user := db.NewUser(gothUser.Email).Signup("", "").SetVerified(true).SetName(gothUser.Name).AddLoggedIn()
	if err := user.Create(db.MainSession, c); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")

}
