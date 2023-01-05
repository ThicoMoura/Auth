package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken  = errors.New("token is invalid")
	ErrExpiredToken  = errors.New("token has expired")
	ErrInvalidAccess = errors.New("access payload is invalid")
)

type Payload struct {
	ID        uuid.UUID `json:"ID"`
	Email     string    `json:"Email"`
	IssuedAt  time.Time `json:"IssuedAt"`
	ExpiresAt time.Time `json:"ExpiresAt"`
}

func NewPayload(email string, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
}

func (payload Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
