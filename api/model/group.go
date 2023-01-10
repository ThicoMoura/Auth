package model

import "github.com/google/uuid"

type NewG struct {
	Name   string   `json:"name"`
	Access []string `json:"access"`
}

func (model NewG) Get(field string) interface{} {
	switch field {
	case "Name":
		return model.Name
	case "Access":
		return model.Access
	default:
		return nil
	}
}

type UpdateG struct {
	ID     string   `json:"-"`
	Name   *string  `json:"name"`
	Access []string `json:"Access"`
	Active *bool    `json:"Active"`
}

func (model UpdateG) Get(field string) interface{} {
	switch field {
	case "ID":
		return uuid.MustParse(model.ID)
	case "Name":
		if model.Name != nil {
			return *model.Name
		}
		return nil
	case "Access":
		return model.Access
	case "Active":
		if model.Active != nil {
			return *model.Active
		}
		return nil
	default:
		return nil
	}
}

type FindG struct {
	Name     string `form:"name"`
	PageID   *int32 `form:"id"`
	PageSize *int32 `form:"size"`
}

func (model FindG) Get(field string) interface{} {
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

type Group struct {
	ID     uuid.UUID `json:"ID"`
	Name   string    `json:"Name"`
	Access []Model   `json:"Access,omitempty"`
	Users  []Model   `json:"Users,omitempty"`
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

func NewGroup(id uuid.UUID, name string, access, users []Model, active bool) Model {
	return &Group{
		ID:     id,
		Name:   name,
		Access: access,
		Users:  users,
		Active: active,
	}
}
