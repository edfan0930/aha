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

	s := storage.NewSession(session)
	email := s.GetEmail(c.Request)
	if email == "" {

		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	verified := s.GetVerified(c.Request)
	if verified == "" || verified == "false" {

		c.Redirect(http.StatusSeeOther, "/login/revalidate")
		return
	}

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")

}
