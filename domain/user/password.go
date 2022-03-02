package user

import (
	"errors"
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"gorm.io/gorm"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type resetPassword struct {
	Old     string `json:"old" form:"old" binding:"required"`
	New     string `json:"new" form:"new" binding:"required"`
	Confirm string `json:"confirm" form:"confirm" binding:"required,eqfield=New"`
}

//ResetPassword
func ResetPassword(c *gin.Context) {

	r := &resetPassword{}
	if err := c.Bind(r); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	//
	if err := VerifyPassword(r.New); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	email := c.Request.Header.Get(storage.StorageKey.Email)

	user, err := db.WhereFirst(db.MainSession, c, db.User{Email: email, Password: r.Old})
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, response.Error("The password you have entered is incorrect."))
			return
		}

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	if err := user.UpdatePassword(db.MainSession, c, r.New); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.Success())
}

func ResetPasswordView(c *gin.Context) {

	c.HTML(http.StatusOK, "resetPassword.html", gin.H{})
}
