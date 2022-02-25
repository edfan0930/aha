package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Dashboard(r *gin.Engine) {

	h := r.Group("/dashboard", VerfySession())

	h.GET("/", func(c *gin.Context) {

		email := c.Request.Header.Get("email")
		name := c.Request.Header.Get("name")

		fmt.Println("header", c.Request.Header)
		/* 	c.JSON(http.StatusOK, struct {
			Email string `json:"email"`
		}{Email: email}) */
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"email": email, "name": name,
		})
	})

	h.GET("/profile", func(c *gin.Context) {

		email := c.Request.Header.Get("email")
		name := c.Request.Header.Get("name")

		c.HTML(http.StatusOK, "profile.html", gin.H{
			"email": email,
			"name":  name,
		})
	})

	h.POST("/logout", func(c *gin.Context) {

	})
}
