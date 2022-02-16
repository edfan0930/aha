package callback

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Email(c *gin.Context) {

	token := c.Query("verify_token")
	if token == "" {

		c.JSON(http.StatusBadRequest, "")
	}

	c.JSON(http.StatusOK, "")
}
