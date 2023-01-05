package model

type Login struct {
	Email string `json:"email" validate:"required,email"`
	Pass  string `json:"pass" validate:"required"`
}

func (model Login) Get(field string) interface{} {
	switch field {
	case "Email":
		return model.Email
	case "Pass":
		return model.Pass
	default:
		return nil
	}
}
