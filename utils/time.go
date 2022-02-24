package utils

import "time"

const (
	DateNow = "2006-01-02"
)

//GetDateNow only date
func GetDateNow() time.Time {

	t, _ := time.Parse(DateNow, time.Now().Format(DateNow))
	return t
}
