package model

import "github.com/google/uuid"

type NewG struct {
	Name   string   `json:"name"`
	Access []string `json:"access"`
}

type UpdateG struct {
	Name   string   `json:"name"`
	Access []string `json:"Access"`
}

type Group struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"Name"`
	Access []Access  `json:"Access,omitempty"`
	Users  []User    `json:"Users,omitempty"`
	Active bool      `json:"Active"`
}

func (model Group) Get(field string) interface{} {
	switch field {
	case "ID":
		return model.ID
	case "Name":
		return model.Name
	case "Access":
		return model.Access
	case "Users":
		return model.Users
	case "Active":
		return model.Active
	default:
		return nil
	}
}
