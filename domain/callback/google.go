package callback

import (
	"fmt"
	"net/http"

	"github.com/edfan0930/aha/common/oauth"

	"github.com/gin-gonic/gin"
)

//Google
func Google(c *gin.Context) {

	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}
	fmt.Println("over query")
	g := oauth.NewGoogleOauth2()
	if err := g.Exchange(c, code); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("over auth")
	if err := g.Request(c); err != nil {
		c.JSON(http.StatusRequestTimeout, err.Error())
		return
	}

	fmt.Println("response", g.Response)
	//	oauth.Exchange(oauth.G)
}
