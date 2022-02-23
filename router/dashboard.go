package router

import "github.com/gin-gonic/gin"

func Dashboard(r *gin.Engine) {

	d := r.Group("/dashboard")
	d.GET("/home", func(c *gin.Context) {

		//c.HTML
	})
}
