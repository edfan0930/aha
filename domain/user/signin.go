package user

type (
	SigninRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
)

//NewSignupRequest
func NewSigninRequest() *SigninRequest {
	return &SigninRequest{}
}
