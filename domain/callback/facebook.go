package callback

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/gin-gonic/gin"
)

func Facebook(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	f := oauth.NewFacebookOauth2()
	if err := f.Exchange(c, code); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := f.Request(c); err != nil {
		c.JSON(http.StatusRequestTimeout, err.Error())
		return
	}

	fmt.Println("response", f.Response)
	c.JSON(http.StatusOK, f.Response)

}
