package validation

import "github.com/go-playground/validator/v10"

//ErrorResponse represents validation error
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

//Validate and return error is any data is wrong
func Validate(data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
