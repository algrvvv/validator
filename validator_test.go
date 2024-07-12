package validator_test

import (
	"fmt"
	"github.com/algrvvv/validator"
	"testing"
)

type mockUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type mockMessages struct{}

func (m mockMessages) Required(field string) string {
	return fmt.Sprintf("Field %s is required", field)
}

func (m mockMessages) Min(field string, min int) string {
	return fmt.Sprintf("The length of the field %s must be at least %d ", field, min)
}

func (m mockMessages) Email(field string) string {
	return fmt.Sprintf("The %s field must match the mail format", field)
}

func TestValidate(t *testing.T) {
	t.Run("should return error if email is empty", func(t *testing.T) {
		err := validator.Validate(mockUser{
			Email:    "",
			Password: "p4ssw0rd",
		})

		if err == nil {
			t.Error("error should not be nil")
		} else {
			t.Log(err)
		}
	})

	t.Run("should return error if has no email", func(t *testing.T) {
		err := validator.Validate(mockUser{
			Password: "p4ssw0rd",
		})

		if err == nil {
			t.Error("error should not be nil")
		} else {
			t.Log(err)
		}
	})

	t.Run("should return error if len(password) < 8", func(t *testing.T) {
		err := validator.Validate(mockUser{
			Email:    "example@example.com",
			Password: "p4sd",
		})

		if err == nil {
			t.Error("error should not be nil")
		} else {
			t.Log(err)
		}
	})

	t.Run("should be return nil if email and password comply with the conditions", func(t *testing.T) {
		err := validator.Validate(mockUser{
			Email:    "example@example.com",
			Password: "p4ssw0rd",
		})

		if err != nil {
			t.Error("error should be nil")
		} else {
			t.Log(err)
		}
	})

	t.Run("test use custom messages", func(t *testing.T) {
		err := validator.ValidateWithMessage(mockUser{}, &mockMessages{})
		if err == nil {
			t.Error("error should not be nil")
		} else {
			t.Log(err)
		}
	})
}
