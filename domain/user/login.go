package user

import (
	"net/http"

	"github.com/edfan0930/aha/domain/response"
	"github.com/gin-gonic/gin"
)

type (
	LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
)

func Login(c *gin.Context) {
	r := &LoginRequest{}
	if err := c.BindJSON(r); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
	}
}
