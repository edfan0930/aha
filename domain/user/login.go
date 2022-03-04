package user

import (
	"errors"
	"net/http"

	"github.com/markbates/goth/gothic"

	"github.com/edfan0930/aha/common/storage"

	"gorm.io/gorm"

	"github.com/edfan0930/aha/db"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type (
	LoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required,email"`
		Password string `json:"password" form:"password" binding:"required"`
	}
)

//Login
func Login(c *gin.Context) {

	r := &LoginRequest{}
	if err := c.Bind(r); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	user, err := db.WhereFirst(db.MainSession, c, db.User{Email: r.Email, Password: r.Password})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	session := storage.NewSession(storage.PassSecure(c.Request))
	if err := session.Login(c.Writer, c.Request, user.Email, user.Name, user.Verified); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	if !user.Verified {

		c.Redirect(http.StatusSeeOther, "/revalidate")
		return
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")
}

//OauthLogin
func OauthLogin(c *gin.Context) {

	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {

		u, err := db.First(db.MainSession, c, user.Email)
		if err != nil {

			c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
			return
		}

		session := storage.NewSession(storage.PassSecure(c.Request))
		if err := session.Login(c.Writer, c.Request, user.Email, u.Name, true); err != nil {

			c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
			return
		}

		c.Redirect(http.StatusSeeOther, "/dashboard/profile")
	}

	gothic.BeginAuthHandler(c.Writer, c.Request)
}
