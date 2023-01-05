package token

import "time"

type JWT struct {
	Key      string
	Duration time.Duration
}
