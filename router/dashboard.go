package router

import (
	"net/http"

	"github.com/edfan0930/aha/db"
	"github.com/edfan0930/aha/domain/dashboard"
	"github.com/edfan0930/aha/domain/response"

	"github.com/gin-gonic/gin"
)

func Dashboard(r *gin.Engine) {

	h := r.Group("/dashboard", VerfySession())

	h.GET("", func(c *gin.Context) {

		dashboard.UserData(c)

		/* 	email := c.Request.Header.Get("email")
		name := c.Request.Header.Get("name")

		fmt.Println("header", c.Request.Header)
		/* 	c.JSON(http.StatusOK, struct {
			Email string `json:"email"`
		}{Email: email}) */
		/*	c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"email": email, "name": name,
		}) */
	})

	h.GET("/profile", func(c *gin.Context) {

		email := c.Request.Header.Get("email")
		name := c.Request.Header.Get("name")

		user, err := db.First(db.MainSession, c, email)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}

		var reset bool
		if user.Password != "" {
			reset = true
		}

		c.HTML(http.StatusOK, "profile.html", gin.H{
			"email":      email,
			"name":       name,
			"couldReset": reset,
		})
	})

}
