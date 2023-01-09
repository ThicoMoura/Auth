package model

import "github.com/google/uuid"

type NewA struct {
	SystemID string   `json:"system_id"`
	Table    string   `json:"table"`
	Types    []string `json:"types"`
}

type UpdateA struct {
	Types []string `json:"types"`
}

type Access struct {
	ID       uuid.UUID `json:"ID"`
	SystemID uuid.UUID `json:"SystemID"`
	Table    string    `json:"Table"`
	Types    []string  `json:"Types"`
}

func (model Access) Get(field string) interface{} {
	switch field {
	case "ID":
		return model.ID
	case "SystemID":
		return model.SystemID
	case "Table":
		return model.Table
	case "Types":
		return model.Types
	default:
		return nil
	}
}

func NewAccess(id, systemID uuid.UUID, table string, types []string) Model {
	return &Access{
		ID:       id,
		SystemID: systemID,
		Table:    table,
		Types:    types,
	}
}
