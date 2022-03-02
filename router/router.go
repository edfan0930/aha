package router

import (
	"net/http"

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

	//user methods
	User(r)

	r.Run(":3000")
}
