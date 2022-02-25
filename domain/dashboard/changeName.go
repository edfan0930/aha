package dashboard

import (
	"net/http"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type (
	changeName struct {
		Name string `json:"name" form:"name" binding:"required"`
	}
)

func ChangeName(c *gin.Context) {

	email := c.Request.Header.Get("eamil")
	if email == "" {

		c.JSON(http.StatusBadRequest, response.Error("your session has been terminated"))
		return
	}

	r := &changeName{}
	if err := c.Bind(r); err != nil {

		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	
}
