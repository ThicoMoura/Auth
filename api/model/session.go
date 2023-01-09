package model

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID        uuid.UUID `json:"ID"`
	UserID    uuid.UUID `json:"UserID"`
	Token     string    `json:"Token"`
	Ip        string    `json:"Ip"`
	Agent     string    `json:"Agent"`
	CreatedAt time.Time `json:"CreatedAt"`
	ExpiresAt time.Time `json:"ExpiresAt"`
}

func (model Session) Get(field string) interface{} {
	switch field {
	case "ID":
		return model.ID
	case "UserID":
		return model.UserID
	case "Token":
		return model.Token
	case "Ip":
		return model.Ip
	case "Agent":
		return model.Agent
	case "CreatedAt":
		return model.CreatedAt
	case "ExpiresAt":
		return model.ExpiresAt
	default:
		return nil
	}
}
