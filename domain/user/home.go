package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/gin-gonic/gin"
)

//Home
func Home(c *gin.Context) {

	//verify session
	session, err := storage.UserHandler(c.Request)
	if err != nil {

		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	//session methods
	s := storage.NewSession(session)
	//get email value
	email := s.GetEmail(c.Request)
	if email == "" {

		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	//get verified value
	verified := s.GetVerified(c.Request)
	if verified == "" || verified == "false" {

		c.Redirect(http.StatusSeeOther, "/revalidate")
		return
	}

	//redirect to profile if login already and pass validate
	c.Redirect(http.StatusSeeOther, "/dashboard/profile")
}
