package callback

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/oauth"
	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

func Facebook(c *gin.Context) {

	code := c.Query("code")
	if code == "" {

		c.JSON(http.StatusBadRequest, response.Error("bad request"))
		return
	}
	fmt.Println("code", code)

	f := oauth.NewFacebookOauth2()
	if err := f.Exchange(c, code); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	fmt.Println("over exchange")
	if err := f.Request(c); err != nil {
		c.JSON(http.StatusRequestTimeout, err.Error())
		return
	}

	fmt.Println("response", f.Response)
	c.JSON(http.StatusOK, f.Response)

}
