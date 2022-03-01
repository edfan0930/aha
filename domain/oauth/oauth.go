package oauth

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/storage"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func Oauth(c *gin.Context) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {

		fmt.Println("in oauth")
		storage.NewSession(storage.PassSecure(c.Request)).Login(c.Writer, c.Request, gothUser.Email, gothUser.Name)

		c.Redirect(http.StatusSeeOther, "/dashboard/profile")
		return
	}
	fmt.Println("---------------------store", gothic.Store)
	gothic.BeginAuthHandler(c.Writer, c.Request)

}
