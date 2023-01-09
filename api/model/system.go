package model

import "github.com/google/uuid"

type NewSy struct {
	Name string `json:"name"`
}

type UpdateSy struct {
	Name string `json:"name"`
}

type System struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"Name"`
	Access []Access  `json:"Access,omitempty"`
	Active bool      `json:"Active"`
}

func (model System) Get(field string) interface{} {
	switch field {
	case "ID":
		return model.ID
	case "Name":
		return model.Name
	case "Access":
		return model.Access
	case "Active":
		return model.Active
	default:
		return nil
	}
}
