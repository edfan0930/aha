package router

import (
	"net/http"

	"github.com/edfan0930/aha/domain/user"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func User(r *gin.Engine) {

	u := r.Group("/user", VerfySession(), Verified())

	//reset password
	//normally it must be "put method"
	u.POST("/password", user.ResetPassword)

	//reset password
	u.GET("/password-view", user.ResetPasswordView)

	//update name
	//normally it must be "put method"
	u.POST("/name", user.ResetName)

	//reset name view
	u.GET("/name-view", func(c *gin.Context) {

		name := c.Request.Header.Get("name")
		c.HTML(http.StatusOK, "updateName.html", gin.H{"name": name})
	})

	//logout
	u.GET("/logout", func(c *gin.Context) {

		gothic.Logout(c.Writer, c.Request)
		user.Logout(c)
	})
}
