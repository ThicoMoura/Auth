package token

import "time"

type Maker interface {
	New(email string, duration time.Duration) (string, *Payload, error)
	Valid(token string) (*Payload, error)
}