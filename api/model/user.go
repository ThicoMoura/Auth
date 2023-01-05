package model

import "github.com/google/uuid"

type User struct {
	GroupID uuid.UUID `json:"GroupID"`
	Email   string    `json:"Email"`
	Name    string    `json:"Name"`
	Pass    string    `json:"-"`
	Active  bool      `json:"Active"`
}

func (model User) Get(field string) interface{} {
	switch field {
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
