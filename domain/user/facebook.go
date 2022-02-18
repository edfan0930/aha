package user

import (
	"net/http"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/gin-gonic/gin"
)

//SignupFacebook
func SignupFacebook(c *gin.Context) {

	redirectURL := oauth.FackbookOAuthURL()
	c.Redirect(http.StatusSeeOther, redirectURL)

}
