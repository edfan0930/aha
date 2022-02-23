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
	SessionAt       = "session_at"
	CreatedAt       = "created_at"
)

const (
	SessionAtFormat = "2006-01-02"
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
		SessionAt   time.Time `gorm:"index"`
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

func First(session *MySQL, context context.Context, email string) (*User, error) {
	u := &User{}
	tx := session.Gorm.First(u, email)
	return u, tx.Error
}

//Signup user defined
func (u *User) Signup(password, verifyToken string) *User {

	u.Password = password
	u.VerifyToken = verifyToken
	u.SessionAt, _ = time.Parse(SessionAtFormat, time.Now().Format(SessionAtFormat))
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

//UpdateLastSession
func (u *User) UpdateLastSession(session *MySQL, context context.Context) error {

	tx := session.Gorm.Begin()
	t := tx.Model(u).Update(Verified, time.Now().Format(SessionAtFormat))
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
