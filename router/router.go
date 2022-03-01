package router

import (
	"net/http"

	"github.com/markbates/goth/gothic"

	"github.com/edfan0930/aha/domain/user"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.Use(requestid.New())

	//main page
	r.GET("/", user.Home)

	//signup view
	r.GET("/signup", func(c *gin.Context) {

		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	//signup post
	r.POST("/signup", user.Signup)

	//Dashboard methods
	Dashboard(r)

	//Callback methods
	Callback(r)

	//login methods
	Login(r)

	u := r.Group("/user")

	u.PUT("/password", user.ResetPassword) //reset password

	u.GET("/logout", func(c *gin.Context) {

		gothic.Logout(c.Writer, c.Request)
		user.Logout(c)
	})

	r.Run(":3000")
}
