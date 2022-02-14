package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

//initRouter initializes router
func InitRouter() {

	r := gin.Default()
	r.Use(requestid.New())
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

		redirectURL := oauth.GoogleOAuthURL()
		c.Redirect(http.StatusSeeOther, redirectURL)
	})

	r.Run(":3000")
}
