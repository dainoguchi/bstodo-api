package testutil

import "time"

func ToStrP(s string) *string {
	return &s
}

func ToTimeP(t time.Time) *time.Time {
	return &t
}
