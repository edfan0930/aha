package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	PasswordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
)

//validate:"password,eqfield=ConfirmPasswords"
type (
	SignupRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,eqfield=ConfirmPassword"`
	}
	SignupConfirm struct {
		SignupRequest
		ConfirmPassword string `json:"confirm_password" binding:"required"`
	}
)

func Signup(c *gin.Context) {
	//	redirectURL := oauth.GoogleOAuthURL()
	//c.Redirect(http.StatusSeeOther, redirectURL)
	r := new(SignupRequest)
	if err := c.ShouldBindJSON(r); err != nil {
		fmt.Println("err", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

//NewSignupRequest
func NewSignupRequest() *SignupRequest {
	return &SignupRequest{}
}

//NewSignupConfirm
func NewSignupConfirm() *SignupConfirm {
	return &SignupConfirm{}
}
