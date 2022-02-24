package callback

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"

	"github.com/edfan0930/aha/domain/response"

	"github.com/edfan0930/aha/utils"

	"github.com/edfan0930/aha/db"

	"github.com/edfan0930/aha/common/oauth"

	"github.com/gin-gonic/gin"
)

//Google
func Google(c *gin.Context) {

	code := c.Query("code")
	if code == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
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

	if err := storage.OauthLogin(c.Writer, c.Request, g.Response.Email); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.Redirect(http.StatusSeeOther, "/home")
}
