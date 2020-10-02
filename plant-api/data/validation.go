package data

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

type Validation struct {
	validation *validator.Validate // refs validator.Validate struct
}

// NewValidation initialize and return the Validation struct
func NewValidation() *Validation {
	validation := validator.New()

	return &Validation{
		validation: validation,
	}
}

type ValidationError struct {
	validator.FieldError // refs validator.FieldError
}

type ValidationErrors []ValidationError

// Validate - Validate the request data from client then return the result
// in the form ValidationErrors
func (validation *Validation) Validate(i interface{}) ValidationErrors {
	var errs validator.ValidationErrors
	validationErrs := validation.validation.Struct(i)

	if validationErrs != nil {
		errs = validationErrs.(validator.ValidationErrors)
	}

	if len(errs) == 0 {
		return nil
	}

	var returnArrs []ValidationError
	for _, err := range errs {
		ve := ValidationError{err.(validator.FieldError)}
		returnArrs = append(returnArrs, ve)
	}

	return returnArrs
}

// Error - Return the formatted validation error string
func (validationError ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		validationError.Namespace(),
		validationError.Field(),
		validationError.Tag(),
	)
}

// Error - Return the list of formatted validation error string
func (validationErrors ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range validationErrors {
		errs = append(errs, err.Error())
	}
	return errs
}
