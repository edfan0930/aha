package user

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/email"
	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/edfan0930/aha/env"

	"github.com/gin-gonic/gin"
)

type (
	verification struct {
		Token   string `json:"token" form:"token" binding:"required"`
		Account string `json:"account" form:"account" binding:"required"`
	}
)

//Verification verification email
func Verification(c *gin.Context) {

	r := new(verification)
	if err := c.BindQuery(r); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := db.NewUser(r.Account)
	if err := user.UpdateVerified(db.MainSession, c, r.Token); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	session := storage.NewSession(storage.PassSecure(c.Request))
	if err := session.Verified(c.Writer, c.Request); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func ResendEmail(c *gin.Context) {

	r := c.Request.Header.Get("email")
	if r == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
	}

	user, err := db.First(db.MainSession, c, r)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
	}

	e := email.NewEmail(r)
	query := fmt.Sprintf("token=%s&account=%s", user.VerifyToken, user.Email)
	e.SetURI(env.ServerDomain+"/verification", query)
	e.VerificationEmail()

	c.JSON(http.StatusOK, response.Success())
}
