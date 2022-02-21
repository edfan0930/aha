package db

const (
	UserTable       = "users"
	Email           = "email"
	Password        = "password"
	ConfirmPassword = "confirm_password"
	Name            = "name"
	Logged          = "logged"
	VerifyToken     = "verify_token"
	Verified        = "verified"
	UpdatedAt       = "updated_at"
	LastSignin      = "last_signin"
	CreatedAt       = "created_at"
)

type (
	User struct {
		Email       string `gorm:"primaryKey"`
		Password    string
		Name        string
		Logged      int64
		VerifyToken string
	}
)
