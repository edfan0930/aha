package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type (
	updateName struct {
		Name string `json:"name" form:"name" binding:"required"`
	}
)

//ResetName
func ResetName(c *gin.Context) {

	email := c.Request.Header.Get("email")
	if email == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
	}

	n := &updateName{}
	if err := c.Bind(n); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	u := db.NewUser(email)
	if err := u.UpdateName(db.MainSession, c, n.Name); err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	session := storage.NewSession(storage.PassSecure(c.Request))
	session.SetValue("name", u.Name).Save(c.Writer, c.Request)

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")
}
