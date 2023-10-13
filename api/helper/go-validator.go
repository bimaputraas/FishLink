package helper

import "github.com/go-playground/validator/v10"

func ValidateStruct(s interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}