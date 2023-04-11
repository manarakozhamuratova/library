package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/manarakozhamuratova/one-lab-task2/internal/model"
)

func ValidateBook(book *model.Book) error {
	// Create a new validator instance
	validate := validator.New()

	// Validate the book struct using the validator
	err := validate.Struct(book)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			// return first validation error message
			return fmt.Errorf("%s validation failed on '%s' tag", e.Field(), e.Tag())
		}
	}

	return nil
}
