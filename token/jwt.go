package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct {
	key string
}

const minKeySize = 32

func NewJWT(key string) (Maker, error) {
	if len(key) < minKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minKeySize)
	}

	return &JWT{key}, nil
}

func (maker JWT) New(email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", nil, err
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwt.SignedString([]byte(maker.key))
	if err != nil {
		return "", nil, nil
	}

	return token, payload, nil
}

func (maker JWT) Valid(token string) (*Payload, error) {
	fn := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.key), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, fn)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
