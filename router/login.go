package router

import (
	"net/http"

	"github.com/edfan0930/aha/domain/user"
	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {

	//
	login := r.Group("/login", SetProvider())

	//帳密登入
	login.POST("", user.Login)

	//user defined
	login.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//重新寄送 email
	login.GET("/revalidate", func(c *gin.Context) {

		
		c.HTML(http.StatusOK, "revalidate.html", gin.H{})
	})

	//驗證mail
	login.GET("/verification", user.Verification)

	//oauth2 login
	login.GET("/:provider", user.OauthLogin)

}
