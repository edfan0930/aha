package router

import (
	"log"
	"net/http"
	"time"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

func example() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

//VerfySession
func VerfySession() gin.HandlerFunc {
	return func(c *gin.Context) {

		email, err := storage.GetEmail(c.Request)
		if err != nil {

			c.JSON(http.StatusUnauthorized, response.Error(err.Error()))
			return
		}

		logged, err := storage.GetLoggedOn(c.Request)
		if err != nil {

			c.JSON(http.StatusUnauthorized, response.Error(err.Error()))
			return
		}
		c.Request.Header.Add("email", email)

		if logged {
			c.Request.Header.Add("logged", "true")
		}

		if !logged {
			c.Request.Header.Add("logged", "false")
		}

		c.Next()
	}
}
