package callback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Google
func Google(c *gin.Context) {

	code := c.Query("code")
	if code == "" {

		c.JSON(http.StatusBadRequest, "")
		return
	}

	//	oauth.Exchange(oauth.G)
}
