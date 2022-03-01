package db

import (
	"context"
	"errors"
	"time"

	"github.com/edfan0930/aha/utils"
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
	SessionAt       = "session_at"
	CreatedAt       = "created_at"
)

type (
	User struct {
		Email       string `gorm:"primaryKey" json:"email"`
		Password    string
		Name        string `json:"name"`
		LoggedIn    int64  `json:"logged_in"`
		VerifyToken string
		Verified    bool      `gorm:"index"`
		UpdatedAt   time.Time `gorm:"index"`
		SessionAt   time.Time `gorm:"index" json:"session_at"`
		CreatedAt   time.Time `json:"created_at"`
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

//First
func First(session *MySQL, context context.Context, email string) (*User, error) {
	u := &User{}
	tx := session.Gorm.First(u, "email = ?", email)
	return u, tx.Error
}

//WhereFirst
func WhereFirst(session *MySQL, context context.Context, u User) (*User, error) {

	user := &User{}
	tx := session.Gorm.Where(&u).First(user)

	return user, tx.Error
}

//Signup user defined
func (u *User) Signup(password, verifyToken string) *User {

	u.Password = password
	u.VerifyToken = verifyToken
	u.SessionAt = utils.GetDateNow()
	return u
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

func (u *User) SetVerified(v bool) *User {
	u.Verified = v
	return u
}

func (u *User) SetName(name string) *User {
	u.Name = name
	return u
}

func (u *User) AddLoggedIn() *User {
	u.LoggedIn++
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
		_, err := First(session, context, u.Email)
		if err != nil {

			return tx.AddError(errors.New("verification failed"))
		}
		return nil
	}

	return tx.Error
}

//UpdateLastSession update
func (u *User) UpdateSessionAt(session *MySQL, context context.Context) error {

	tx := session.Gorm.Begin()
	t := tx.Model(u).Update(SessionAt, utils.GetDateNow())
	if t.RowsAffected == 0 {
		_, err := First(session, context, u.Email)
		if err != nil {

			t.AddError(err)
		}

		tx.Rollback()
		return t.Error
	}

	tx.Commit()
	return t.Error
}

//UpdateName
func (u *User) UpdateName(session *MySQL, context context.Context, name string) error {

	tx := session.Gorm.Begin()
	t := tx.Model(u).Update(Name, name)
	if t.RowsAffected == 0 {
		_, err := First(session, context, u.Email)
		if err != nil {

			t.AddError(err)
		}

		tx.Rollback()
		return t.Error
	}

	tx.Commit()
	return t.Error
}

//UpdatePassword
func (u *User) UpdatePassword(session *MySQL, context context.Context, password string) error {

	tx := session.Gorm.Begin()
	t := tx.Model(u).Update(Password, password)
	if t.RowsAffected == 0 {
		_, err := First(session, context, u.Email)
		if err != nil {

			t.AddError(err)
		}

		tx.Rollback()
		return t.Error
	}

	tx.Commit()
	return t.Error
}

func (u *User) Save(session *MySQL, context context.Context) error {

	tx := session.Gorm.Save(u)
	return tx.Error
}
