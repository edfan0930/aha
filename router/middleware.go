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

		s, err := storage.UserHandler(c.Request)
		if err != nil {

			c.JSON(http.StatusUnauthorized, response.Error(err.Error()))
			return
		}

		session := storage.NewSession(s)
		email := session.GetEmail(c.Request)
		logged := session.GetLoggedOn(c.Request)
		name := session.GetName(c.Request)

		if logged == "" {
			logged = "false"
		}

		c.Request.Header.Add("logged", logged)
		c.Request.Header.Add("email", email)
		c.Request.Header.Add("name", name)

		c.Next()
	}
}

func Verified() gin.HandlerFunc {
	return func(c *gin.Context) {

		s := storage.NewSession(storage.PassSecure(c.Request))
		isVerified := s.GetVerified(c.Request)
		if isVerified != "true" {
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		c.Next()
	}
}

func HasLogged() gin.HandlerFunc {
	return func(c *gin.Context) {

		s, err := storage.UserHandler(c.Request)
		if err == nil {

			logged := storage.NewSession(s).GetLoggedOn(c.Request)
			if logged == "true" {

				c.Redirect(http.StatusSeeOther, "/")
				return
			}
		}

		c.Next()
	}
}

//SetProvider
func SetProvider() gin.HandlerFunc {
	return func(c *gin.Context) {

		p := c.Param("provider")
		if p != "" {
			v := c.Request.URL.Query()
			v.Set("provider", p)
			c.Request.URL.RawQuery = v.Encode()
		}

		c.Next()
	}
}
