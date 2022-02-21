package router

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/domain/callback"

	"github.com/gorilla/sessions"

	"github.com/edfan0930/aha/common/email"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/edfan0930/aha/common/storage"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.Use(requestid.New())

	u := r.Group("/user")
	u.GET("/google", func(c *gin.Context) {

		redirectURL := oauth.GoogleOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})

	c := r.Group("/callback")
	c.GET("google", callback.Google)
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
	r.GET("/facebook", func(c *gin.Context) {
		redirectURL := oauth.FackbookOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})
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

		fmt.Println("path", c.Request.URL)
		c.Request.URL.Path = "/session/new"
		r.HandleContext(c)
		return
		c.Redirect(http.StatusSeeOther, "http://localhost:3000/session/new")
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

	r.GET("/session/get", func(c *gin.Context) {
		store, err := storage.Store.Get(c.Request, "user")
		if err != nil {
			fmt.Println("err", err)
		}
		fmt.Printf("store:%s", store)
	})

	r.GET("/session/new", func(c *gin.Context) {
		storage.Store = sessions.NewCookieStore([]byte(storage.GenerSessionID()))
		c.JSON(http.StatusOK, "hello world")
	})

	r.GET("/session/set", func(c *gin.Context) {
		//		c.Cookie
		store, _ := storage.Store.Get(c.Request, "user")
		store.Values["age"] = 18
		err := store.Save(c.Request, c.Writer)
		if err != nil {
			return
		}
		fmt.Println("hello world")
	})

	r.GET("/email", func(c *gin.Context) {
		email.VerificationEmail([]string{""})
	})

	signup := u.Group("/signup")
	signup.GET("", func(c *gin.Context) {
		fmt.Println("signup")
	})

	signup.GET("/google", func(c *gin.Context) {})

	r.Run(":3000")
}
