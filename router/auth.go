package router

import (
	"github.com/gin-gonic/gin"
)

func Auth2(r *gin.Engine) {

	a := r.Group("/oauth", SetProvider())

	a.GET("/:provider", func(c *gin.Context) {})
}
