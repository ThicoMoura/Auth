package model

import "github.com/google/uuid"

type Model interface {
	Get(string) interface{}
}

type Id struct {
	ID string `uri:"id"`
}

func (model Id) Get(field string) interface{} {
	switch field {
	case "ID":
		return uuid.MustParse(model.ID)
	default:
		return nil
	}
}

type ModelA struct {
	Type   string
	Access []uuid.UUID `json:"access"`
}

func (model ModelA) Get(field string) interface{} {
	switch field {
	case "Type":
		return model.Type
	case "Access":
		return model.Access
	default:
		return nil
	}
}
