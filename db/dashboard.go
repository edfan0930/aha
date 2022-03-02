package db

import (
	"context"
	"time"

	"github.com/edfan0930/aha/utils"
)

type (
	userStatistics struct {
		//		LastSeven
		Total         int `json:"total" gorm:"total"`
		SessionsToday int `json:"sessions_today"`
		ActiveAVG     int `json:"active_avg"`
	}

	LastSeven struct {
		ActiveSessions []int   `json:"active_sessions"`
		ActiveAVG      float64 `json:"active_avg"`
	}

	ActiveSession struct {
		Date  string `json:"date"`
		Total int    `json:"total"`
	}

	Users []User
)

//SevenDaySession active session users in the last 7 days
func SevenDaySession(session *MySQL, context context.Context) error {

	return nil
}

//SevenDayRange last 7 days range
func SevenDayRange() (start, end string) {

	end = time.Now().Format("2006-01-02")

	start = time.Now().AddDate(0, 0, -6).Format("2006-01-02")

	return
}

//TotalSignedUp
func TotalSignedUp() (int, error) {

	return 0, nil
}

func UserStatistics() (*userStatistics, error) {

	u := &userStatistics{}
	tx := MainSession.Gorm.Raw("SELECT COUNT(*) AS total FROM users").Scan(u)
	if tx.Error != nil {

		return nil, tx.Error
	}

	tx = MainSession.Gorm.Raw("SELECT COUNT(*) AS sessions_today FROM users WHERE session_at = ?", utils.GetDateNow()).Scan(u)
	if tx.Error != nil {

		return nil, tx.Error
	}

	start, end := SevenDayRange()
	l := new(LastSeven)
	tx = MainSession.Gorm.Raw("SELECT COUNT(*) AS active_sessions FROM users WHERE session_at BETWEEN ? AND ? GROUP BY session_at", start, end).Scan(&l.ActiveSessions)
	if tx.Error != nil {

		return nil, tx.Error
	}

	var total int
	for _, v := range l.ActiveSessions {
		total += v
	}

	u.ActiveAVG = total / len(l.ActiveSessions)

	return u, nil
}

//UserList
func UserList() (*Users, error) {

	u := new(Users)

	//	tx := MainSession.Gorm.Raw("SELECT email,name,logged_in,session_at,created_at FROM users").Scan(u)
	tx := MainSession.Gorm.Model(&User{}).Select("email", "name", "logged_in", "session_at", "created_at").Find(u)
	return u, tx.Error
}
