package user

import (
	"errors"
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type resetPassword struct {
	Old     string `json:"old" form:"old" binding:"required"`
	New     string `json:"new" form:"new" binding:"required"`
	ReEnter string `json:"re_enter" form:"re_enter" binding:"required,eqfield=New"`
}

//ResetPassword
func ResetPassword(c *gin.Context) {

	r := &resetPassword{}
	if err := c.BindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	if err := VerifyPassword(r.New); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	email, err := storage.GetEmail(c.Request)
	if err != nil || email == "" {
		if email == "" {
			err = errors.New("data not found")
		}
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	//
	if err := db.NewUser(email).UpdatePassword(db.MainSession, c, r.New); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
}
