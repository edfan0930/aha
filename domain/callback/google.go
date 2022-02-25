package callback

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/storage"

	"github.com/edfan0930/aha/domain/response"

	"github.com/edfan0930/aha/utils"

	"github.com/edfan0930/aha/db"

	"github.com/edfan0930/aha/common/oauth"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

//Google
func Google(c *gin.Context) {
	c.Request.Header.Set("provider", "Google")
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	fmt.Println("user", user)

	return

	code := c.Query("code")
	if code == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
	}

	state := c.Query("state")
	if state != oauth.UUID {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
	}

	g := oauth.NewGoogleOauth2()
	if err := g.Exchange(c, code); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := g.Request(c); err != nil {
		c.JSON(http.StatusRequestTimeout, response.Error(err.Error()))
		return
	}

	u := db.NewUser(g.Response.Email)
	u.Verified = true
	u.SessionAt = utils.GetDateNow()
	u.Name = g.Response.Name
	u.LoggedIn = 1

	tx := db.MainSession.Gorm.FirstOrCreate(&db.User{}, u)
	if tx.Error != nil {

		c.JSON(http.StatusInternalServerError, response.Error(tx.Error.Error()))
		return
	}

	s := storage.NewSession(storage.PassSecure(c.Request))
	if err := s.Login(c.Writer, c.Request, u.Email, u.Name); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}
