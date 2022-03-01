package callback

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/storage"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

//Oauth2
func Oauth2(c *gin.Context) {

	gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {

		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}

	fmt.Println("user", gothUser)
	storage.NewSession(storage.PassSecure(c.Request)).Login(c.Writer, c.Request, gothUser.Email, gothUser.Name, true)

	c.Redirect(http.StatusSeeOther, "/dashboard/profile")

}
