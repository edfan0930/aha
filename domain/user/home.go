package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {

	session, err := storage.UserHandler(c.Request)
	if err != nil {

		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	l := session.Values[storage.StorageKey.Email]
	email, _ := l.(string)
	if email == "" {

		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	v := session.Values[storage.StorageKey.Verified]
	verified, _ := v.(string)
	if verified == "" || verified == "false" {

		c.Redirect(http.StatusSeeOther, "/login/revalidate")
		return
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")

}
