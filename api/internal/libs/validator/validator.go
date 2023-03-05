package validator

import (
	"reflect"
	"strings"

	v "github.com/go-playground/validator/v10"
)

// Error response
type errorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// Validates a struct
func ValidateStruct(s interface{}) []*errorResponse {
	var errors []*errorResponse
	validate := v.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(v.ValidationErrors) {
			errors = append(errors, &errorResponse{
				Field: strings.ToLower(err.Field()),
				Tag:   err.Tag(),
				Param: err.Param(),
			})
		}
	}

	return errors
}
