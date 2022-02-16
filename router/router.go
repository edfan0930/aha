package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/domain/user"

	"github.com/edfan0930/aha/common/email"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.Use(requestid.New())

	u := r.Group("/user")
	u.POST("/signup", user.Signup)
	r.GET("/callback", func(c *gin.Context) {
		s := c.Query("state")
		fmt.Println("state:", s)
		//		code:=c.Query("code")
		c.JSON(http.StatusOK, "isOk")
	})

	r.GET("/auth", func(c *gin.Context) {
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

	r.GET("/google", func(c *gin.Context) {

		redirectURL := oauth.GoogleOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})

	r.GET("/facebook", func(c *gin.Context) {
		redirectURL := oauth.FackbookOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})

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
	})

	r.GET("/main", func(c *gin.Context) {

		code := c.Query("code")
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

		//		c.JSON(http.StatusOK, c.Query("code"))
	})

	r.GET("/email", func(c *gin.Context) {
		email.VerificationEmail([]string{""})
	})

	r.Run(":3000")
}
