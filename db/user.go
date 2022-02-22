package db

import (
	"context"
	"errors"
	"time"
)

const (
	UserTable       = "users"
	Email           = "email"
	Password        = "password"
	ConfirmPassword = "confirm_password"
	Name            = "name"
	LoggedIn        = "logged_in" //Number of times logged in.
	VerifyToken     = "verify_token"
	Verified        = "verified"
	UpdatedAt       = "updated_at"
	LastSession     = "last_session"
	CreatedAt       = "created_at"
)

type (
	User struct {
		Email       string `gorm:"primaryKey"`
		Password    string
		Name        string
		LoggedIn    int64
		VerifyToken string
		Verified    bool      `gorm:"index"`
		UpdatedAt   time.Time `gorm:"index"`
		LastSession time.Time `gorm:"index"`
		CreatedAt   time.Time
	}
)

//AutoMigrate
func AutoMigrate(session *MySQL) error {

	return session.Gorm.AutoMigrate(&User{})
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}

//SetPassword
func (u *User) SetPassword(p string) *User {
	u.Password = p
	return u
}

//SetVerifyToken
func (u *User) SetVerifyToken(v string) *User {
	u.VerifyToken = v
	return u
}

//Signup
func (u *User) Signup(password, verifyToken string) *User {
	u.Password = password
	u.VerifyToken = verifyToken
	return u
}

func (u *User) Create(session *MySQL, context context.Context) error {

	result := session.Gorm.Create(u)
	return result.Error
}

//UpdateVerified
func (u *User) UpdateVerified(session *MySQL, context context.Context, token string) error {

	tx := session.Gorm.Model(u).Where(VerifyToken+" = ?", token).Update(Verified, true)
	if tx.RowsAffected == 0 {

		return tx.AddError(errors.New("verification failed"))
	}
	return tx.Error
}
