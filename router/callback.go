package router

import (
	"github.com/edfan0930/aha/domain/callback"
	"github.com/gin-gonic/gin"
)

func Callback(r *gin.Engine) {

	c := r.Group("/callback", SetProvider())

	c.GET("/:provider", callback.Oauth2)

	//	c.GET("google", callback.Google)

	//c.GET("facebook", callback.Facebook)

}
