package router

import (
	"github.com/edfan0930/aha/domain/dashboard"

	"github.com/gin-gonic/gin"
)

func Dashboard(r *gin.Engine) {

	h := r.Group("/dashboard", VerfySession(), Verified())

	h.GET("", dashboard.UserData)

	h.GET("/profile", dashboard.Profile)

}
