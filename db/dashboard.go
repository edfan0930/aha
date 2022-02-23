package db

import (
	"context"
	"time"
)

type (
	LastSeven struct {
		ActiveSession []ActiveSession `json:"active_session"`
		AVG           int             `json:"avg"`
	}

	ActiveSession struct {
		Date  string `json:"date"`
		Total int    `json:"total"`
	}
)

//SevenDaySession active session users in the last 7 days
func SevenDaySession(session *MySQL, context context.Context) error {

	return nil
}

//SevenDayRange last 7 days range
func SevenDayRange() (start, end string) {

	end = time.Now().Format("2006-01-02")

	start = time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	return
}
