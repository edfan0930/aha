package router

import (
	"net/http"

	"github.com/edfan0930/aha/domain/user"
	"github.com/gin-gonic/gin"
)

func Login(r *gin.Engine) {

	//
	login := r.Group("/login", SetProvider(), HasLogged())

	//帳密登入
	login.POST("", user.Login)

	//login view
	login.GET("", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	//驗證mail
	//	login.GET("/verification", user.Verification)

	//oauth2 login
	login.GET("/:provider", user.OauthLogin)
}
