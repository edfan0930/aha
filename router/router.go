package router

import (
	"net/http"

	"github.com/edfan0930/aha/domain/user"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/edfan0930/aha/common/storage"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
			"title": "test",
		})
	})
	r.Use(requestid.New())

	//Dashboard methods
	Dashboard(r)

	//Callback methods
	Callback(r)

	//
	r.GET("/", func(c *gin.Context) {
		session, err := storage.UserHandler(c.Request)
		if err != nil {

			c.Redirect(http.StatusSeeOther, "/signup")
			return
		}

		l := session.Values[storage.StorageKey.Email]
		email, _ := l.(string)
		if email == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}

		c.Redirect(http.StatusSeeOther, "/dashboard")
	})

	u := r.Group("/user")

	u.POST("/login", user.Login) //user defined login

	u.PUT("/password", user.ResetPassword) //reset password

	/*
		//	u.POST("/signup", user.Signup)
		r.GET("/callback", func(c *gin.Context) {
			s := c.Query("state")
			fmt.Println("state:", s)
			//		code:=c.Query("code")
			c.JSON(http.StatusOK, "isOk")
		}) */

	/* r.GET("/auth", func(c *gin.Context) {
		if code := c.Query("code"); code != "" {
			fmt.Println("get code")
			token, err := oauth.GoogleExchange(context.Background(), code)
			if err != nil {
				c.JSON(http.StatusOK, err.Error())
				return
			}
			if token.AccessToken != "" {
				oauth.GoogleClient(context.Background(), token)
			}
			fmt.Println("AccessToken empty")
			c.JSON(http.StatusOK, token)
			return
		}

	})
	*/

	/*
		r.GET("/facebook/callback", func(c *gin.Context) {
			code := c.Query("code")
			token, err := oauth.FacebookExchange(context.Background(), code)
			if err != nil {
				c.JSON(http.StatusOK, err.Error())
				return
			}
			if token.AccessToken != "" {
				oauth.GoogleClient(context.Background(), token)
			}
			fmt.Println("AccessToken empty")
			c.JSON(http.StatusOK, token)
		}) */

	r.GET("/main", func(c *gin.Context) {

		c.Redirect(http.StatusSeeOther, "/")
		/* 		code := c.Query("code")
		   		token, err := oauth.FacebookExchange(context.Background(), code)
		   		if err != nil {
		   			c.JSON(http.StatusOK, err.Error())
		   			return
		   		}
		   		if token.AccessToken != "" {
		   			oauth.FacebookClient(context.Background(), token)
		   			return
		   		}
		   		fmt.Println("AccessToken empty")
		   		c.JSON(http.StatusOK, token)
		*/
		//		c.JSON(http.StatusOK, c.Query("code"))
	})

	signup := r.Group("/signup")

	signup.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	signup.POST("", user.Signup)

	signup.GET("/verification", user.Verification)

	login := r.Group("/login")

	login.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	login.GET("/google", func(c *gin.Context) {

		redirectURL := oauth.GoogleOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})

	login.GET("/facebook", func(c *gin.Context) {
		redirectURL := oauth.FackbookOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)

	})

	r.Run(":3000")
}
