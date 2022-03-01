package router

import (
	"net/http"

	"github.com/markbates/goth/gothic"

	"github.com/edfan0930/aha/domain/user"

	"github.com/edfan0930/aha/common/storage"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/index", func(c *gin.Context) {

		text := []string{"a", "b", "c", "d"}
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"list": text,
		})
	})

	r.Use(requestid.New())

	//
	r.GET("/", func(c *gin.Context) {

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

		c.Redirect(http.StatusSeeOther, "/profile")
	})

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
	//

	u := r.Group("/user")

	//	u.POST("/login", user.Login) //user defined login

	u.PUT("/password", user.ResetPassword) //reset password
	logout := u.Group("/logout", SetProvider())
	logout.GET("", user.Logout)
	logout.GET(":provider", func(c *gin.Context) {

		gothic.Logout(c.Writer, c.Request)
		c.Redirect(http.StatusSeeOther, "/")
	})
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

	/* 	login.GET("/facebook", func(c *gin.Context) {

		redirectURL := oauth.FackbookOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	}) */

	r.Run(":3000")
}
