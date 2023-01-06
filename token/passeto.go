package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type Paseto struct {
	paseto *paseto.V2
	key    []byte
}

func NewPaseto(key string) (Maker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exaclty %d characteres", chacha20poly1305.KeySize)
	}

	return &Paseto{
		paseto: paseto.NewV2(),
		key:    []byte(key),
	}, nil
}

func (maker Paseto) New(email string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(email, duration)
	if err != nil {
		return "", nil, err
	}

	token, err := maker.paseto.Encrypt(maker.key, payload, nil)
	if err != nil {
		return "", nil, err
	}

	return token, payload, nil
}

func (maker Paseto) Valid(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.key, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
