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
	//put
	u.POST("/password", user.ResetPassword)

	u.GET("/password", func(c *gin.Context) {

	})

	//update name
	//put
	u.POST("/name", user.ResetName)

	//update name view
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
