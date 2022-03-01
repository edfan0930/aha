package db

import (
	"context"
	"time"

	"github.com/edfan0930/aha/utils"
)

type (
	userStatistics struct {
		//		LastSeven
		Total        int `json:"total" gorm:"total"`
		SessionToday int `json:"session_today"`
		ActiveAVG    int `json:"active_avg"`
	}

	LastSeven struct {
		ActiveSession []int   `json:"active_session"`
		ActiveAVG     float64 `json:"active_avg"`
	}

	ActiveSession struct {
		Date  string `json:"date"`
		Total int    `json:"total"`
	}

	Users struct {
	}
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

func SessionToday() (*Users, error) {

	return &Users{}, nil
}

func UserStatistics() (*userStatistics, error) {

	u := &userStatistics{}
	tx := MainSession.Gorm.Raw("SELECT COUNT(*) AS total FROM users").Scan(u)
	if tx.Error != nil {

		return nil, tx.Error
	}

	tx = MainSession.Gorm.Raw("SELECT COUNT(*) AS session_today FROM users WHERE session_at = ?", utils.GetDateNow()).Scan(u)
	if tx.Error != nil {

		return nil, tx.Error
	}

	start, end := SevenDayRange()
	l := new(LastSeven)
	tx = MainSession.Gorm.Raw("SELECT COUNT(*) AS active_session FROM users WHERE session_at BETWEEN ? AND ? GROUP BY session_at", start, end).Scan(&l.ActiveSession)
	if tx.Error != nil {

		return nil, tx.Error
	}

	var total int
	for _, v := range l.ActiveSession {
		total += v
	}

	u.ActiveAVG = total / len(l.ActiveSession)

	return u, nil
}
