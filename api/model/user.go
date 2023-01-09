package model

import "github.com/google/uuid"

type NewU struct {
	GroupID string   `json:"group_id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Pass    string   `json:"pass"`
	Access  []string `json:"access"`
}

type UpdateU struct {
	GroupID string `json:"group_id"`
	Name    string `json:"name"`
	Pass    string `json:"pass"`
}

type User struct {
	ID      uuid.UUID `json:"ID"`
	GroupID uuid.UUID `json:"GroupID"`
	Email   string    `json:"Email"`
	Name    string    `json:"Name"`
	Pass    string    `json:"-"`
	Access  []Access  `json:"Access,omitempty"`
	Active  bool      `json:"Active"`
}

func (model User) Get(field string) interface{} {
	switch field {
	case "D":
		return model.ID
	case "GroupID":
		return model.GroupID
	case "Email":
		return model.Email
	case "Name":
		return model.Name
	case "Pass":
		return model.Pass
	case "Active":
		return model.Active
	default:
		return nil
	}
}
