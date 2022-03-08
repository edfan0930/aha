package callback

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"

	"github.com/edfan0930/aha/domain/response"

	"github.com/gin-gonic/gin"
)

func Email(c *gin.Context) {

	token := c.Query("verify_token")
	if token == "" {
		c.JSON(http.StatusBadRequest, response.Error("verify_token is required"))
		return
	}

	email := c.Query("email")
	if email == "" {

		c.JSON(http.StatusBadRequest, response.Error("email is required"))
		return
	}

	user, err := db.First(db.MainSession, c, email)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	if user.VerifyToken != token {

		c.JSON(http.StatusBadRequest, response.Error("wrong token"))
		return
	}

	if !user.Verified {
		user.Verified = true
		if err := user.Save(db.MainSession, c); err != nil {
			c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
			return
		}
	}

	s := storage.NewSession(storage.PassSecure(c.Request))
	s.Verified(c.Writer, c.Request)

	c.Redirect(http.StatusSeeOther, "/")
}
