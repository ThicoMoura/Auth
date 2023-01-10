package model

import "github.com/google/uuid"

type NewSy struct {
	Name string `json:"name"`
}

func (model NewSy) Get(field string) interface{} {
	switch field {
	case "Name":
		return model.Name
	default:
		return nil
	}
}

type UpdateSy struct {
	ID     string  `uri:"id"`
	Name   *string `json:"name"`
	Active *bool   `json:"active"`
}

func (model UpdateSy) Get(field string) interface{} {
	switch field {
	case "ID":
		return uuid.MustParse(model.ID)
	case "Name":
		if model.Name != nil {
			return *model.Name
		}
		return nil
	case "Active":
		if model.Active != nil {
			return *model.Active
		}
		return nil
	default:
		return nil
	}
}

type FindSy struct {
	Name     string `form:"name"`
	PageID   *int32 `form:"id"`
	PageSize *int32 `form:"size"`
}

func (model FindSy) Get(field string) interface{} {
	switch field {
	case "Name":
		return model.Name + "%"
	case "PageID":
		return model.PageID
	case "PageSize":
		return model.PageSize
	case "Limit":
		if model.PageSize != nil {
			return *model.PageSize
		}
		return nil
	case "Offset":
		if model.PageSize != nil && model.PageID != nil {
			return int32(*model.PageSize * (*model.PageID - 1))
		}
		return nil
	default:
		return nil
	}
}

type System struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"Name"`
	Access []Model   `json:"Access,omitempty"`
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

func NewSystem(id uuid.UUID, name string, access []Model, active bool) Model {
	return &System{
		ID:     id,
		Name:   name,
		Access: access,
		Active: active,
	}
}
