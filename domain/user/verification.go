package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/email"
	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"

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

	//get query value
	r := new(verification)
	if err := c.BindQuery(r); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	//update user to verified
	user := db.NewUser(r.Account)
	if err := user.UpdateVerified(db.MainSession, c, r.Token); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	//update session storage
	session := storage.NewSession(storage.PassSecure(c.Request))
	if err := session.Verified(c.Writer, c.Request); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//redirect page
	c.Redirect(http.StatusSeeOther, "/")
}

//ResendEmail resend email
func ResendEmail(c *gin.Context) {

	//get user email from session storage
	r := c.Request.Header.Get("email")
	if r == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
	}

	//get user query
	user, err := db.First(db.MainSession, c, r)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
	}

	//email
	e := email.NewEmail(r)
	e.SetQuery(user.VerifyToken, user.Email).SetURI()
	if err := e.VerificationEmail(); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success())
}
