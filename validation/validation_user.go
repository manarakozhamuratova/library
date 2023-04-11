package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

func ValidateUser(user *model.User) error {
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		// iterate through validation errors
		for _, e := range err.(validator.ValidationErrors) {
			// return first validation error message
			return fmt.Errorf("%s validation failed on '%s' tag", e.Field(), e.Tag())
		}
	}
	return nil
}
