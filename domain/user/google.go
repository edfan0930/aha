package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/edfan0930/aha/common/oauth"
)

//SignupGoogle
func SignupGoogle(c *gin.Context) {

	redirectURL := oauth.GoogleOAuthURL()
	c.Redirect(http.StatusSeeOther, redirectURL)
}
