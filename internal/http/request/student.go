package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

const SBUStudentIDLen = 8

// nolint: tagliatelle
type Student struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s Student) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required, validation.Length(SBUStudentIDLen, SBUStudentIDLen), is.Digit),
		validation.Field(&s.FirstName, validation.Required, is.UTFLetterNumeric),
		validation.Field(&s.LastName, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("student validation failed %w", err)
	}

	return nil
}
