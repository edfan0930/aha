package router

import (
	"net/http"

	"github.com/edfan0930/aha/domain/user"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

const (
	PathSignup        = "/signup"
	PathLogin         = "/login"
	PathVerfication   = "/login/verification"
	PathOauthLogin    = "/login/:provider"
	PathDashBoard     = "/dashboard"
	PathProfile       = "/dashboard/profile"
	PathOauthCallback = "/callback/:provider"
	PathUser          = "/user"
	PathPassword      = "/user/password"
	PathName          = "/user/name"
	PathLogout        = "/user/logout"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()

	r.LoadHTMLGlob("view/*")

	r.Use(requestid.New())

	//main page
	r.GET("/", user.Home)

	//驗證mail
	r.GET("/verification", user.Verification)

	signup := r.Group("/signup", HasLogged())
	//signup view
	signup.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	//signup post
	signup.POST("", user.Signup)

	re := r.Group("/revalidate", VerfySession())
	re.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "revalidate.html", gin.H{})
	})

	//resend resend email
	re.GET("/resend", user.ResendEmail)

	//Dashboard methods
	Dashboard(r)

	//Callback methods
	Callback(r)

	//login methods
	Login(r)

	//user methods
	User(r)

	r.Run(":8080")
}
