package model

type Value interface {
	Get(field string) interface{}
}

type model struct {
	fields map[string]interface{}
}

func (model model) Get(field string) interface{} {
	return model.fields[field]
}

func New(fields map[string]interface{}) Value {
	return &model{
		fields: fields,
	}
}
