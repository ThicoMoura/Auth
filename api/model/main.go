package model

import "github.com/google/uuid"

type Model interface {
	Get(string) interface{}
}

type List struct {
	PageID   *int32 `form:"id"`
	PageSize *int32 `form:"size"`
}

func (model List) Get(field string) interface{} {
	switch field {
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
