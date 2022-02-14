package domain

import (
	"net/http"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/gin-gonic/gin"
)

const (
	PasswordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

type (
	signup struct {
	}
)

func Signup(c *gin.Context) {
	redirectURL := oauth.GoogleOAuthURL()
	c.Redirect(http.StatusSeeOther, redirectURL)
}
