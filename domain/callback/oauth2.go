package callback

import (
	"errors"
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

//Oauth2
func Oauth2(c *gin.Context) {

	//oauth2 callback methods
	//get info
	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//save session storage
	if err := storage.NewSession(storage.PassSecure(c.Request)).Login(c.Writer, c.Request, gothUser.Email, gothUser.Name, true); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//get user data
	user, err := db.First(db.MainSession, c, gothUser.Email)
	if err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {

			c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
			return
		}

		//insert row when the record not found
		user := db.NewUser(gothUser.Email).Signup("", "").SetVerified(true).SetName(gothUser.Name).AddLoggedIn()
		if err := user.Create(db.MainSession, c); err != nil {

			c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
			return
		}

		c.Redirect(http.StatusSeeOther, "/dashboard/profile")
		return
	}

	//count user logged
	if err := user.AddLoggedIn().Save(db.MainSession, c); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//redirect to profile dashboard
	c.Redirect(http.StatusSeeOther, "/dashboard/profile")
}
