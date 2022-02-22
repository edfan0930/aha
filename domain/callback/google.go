package callback

import (
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
	g := oauth.NewGoogleOauth2()
	if err := g.Exchange(c, code); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := g.Request(c); err != nil {
		c.JSON(http.StatusRequestTimeout, err.Error())
		return
	}

	c.JSON(http.StatusOK, g.Response)
	//	oauth.Exchange(oauth.G)
}
