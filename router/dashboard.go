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

		fmt.Println("header", c.Request.Header)
		c.JSON(http.StatusOK, struct {
			Email string `json:"email"`
		}{Email: email})
		//c.HTML
	})

}
