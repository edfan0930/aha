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

	signup := r.Group("/signup", HasLogged())
	//signup view
	signup.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	//signup post
	signup.POST("", user.Signup)

	//Dashboard methods
	Dashboard(r)

	//Callback methods
	Callback(r)

	//login methods
	Login(r)

	u := r.Group("/user", VerfySession(), Verified())

	//reset password
	u.PUT("/password", user.ResetPassword)

	//update name
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

	r.Run(":3000")
}
