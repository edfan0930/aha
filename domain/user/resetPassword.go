package user

import (
	"net/http"

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

	if err:= db.NewUser()
}
