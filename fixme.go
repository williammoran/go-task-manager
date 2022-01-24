package tasks

import (
	"strings"
	"time"
)

// FixMe accepts a date in the format YYYY-MM-DD proceeded
// by an arbitrary string. If the current date is after
// the specified date, FixMe will panic. Otherwise it
// does nothing.
func FixMe(msg string) {
	parts := strings.Split(msg, " ")
	deadline, err := time.Parse("2006-01-02", parts[0])
	if err != nil {
		panic(err.Error())
	}
	if time.Now().After(deadline) {
		panic(msg)
	}
}
