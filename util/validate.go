package util

import (
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

type validate struct {
	v *validator.Validate
}

func (validate validate) Validator(s interface{}) map[string]string {
	err := validate.v.Struct(s)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return nil
		}

		errMsg := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			jsonKey := err.Field()
			var message string
			if err.Tag() == "required" {
				message = err.Tag()
			} else {
				message = fmt.Sprintf("'%v' is not in %v format", err.Value(), err.ActualTag())
			}
			errMsg[jsonKey] = message
		}

		return errMsg
	}

	return nil
}

func NewValidate() *validate {
	validate := &validate{
		v: validator.New(),
	}

	validate.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return validate
}
