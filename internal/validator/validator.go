package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationErrors = validator.ValidationErrors

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	RegisterCustomValidations(validate)
}

func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "min_length":
			errors[e.Field()] = fmt.Sprintf("%s must be at least %s characters long", e.Field(), e.Param())
		case "email":
			errors[e.Field()] = "Invalid email format"
		default:
			errors[e.Field()] = "Invalid value"
		}
	}

	return errors
}

func RegisterCustomValidations(v *validator.Validate) {
	validations := map[string]validator.Func{
		"min_length":   MinLength,
		"email_custom": Email,
	}

	for name, fn := range validations {
		if err := v.RegisterValidation(name, fn); err != nil {
			println("failed to register validation %q: %v", name, err)
		}
	}
}

// MinLength Ensures the field's string length is at least N characters long.
func MinLength(fl validator.FieldLevel) bool {
	param := fl.Param()
	minLen, err := strconv.Atoi(param)
	if err != nil {
		return false
	}

	value := fl.Field().String()
	return len(value) >= minLen
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func Email(fl validator.FieldLevel) bool {
	value := strings.TrimSpace(strings.ToLower(fl.Field().String()))
	return emailRegex.MatchString(value)
}
