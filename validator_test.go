package validator_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/algrvvv/validator"
)

type mockUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=12"`
}

type mockWrongUser struct {
	Email    string `json:"email" validate:"required,email,max=l0l"`
	Password string `json:"password" validate:"required,min=something"`
}

type mockProduct struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required,in:1-2-something"`
}

type mockMessages struct{}

func (m mockMessages) Required(field string) string {
	return fmt.Sprintf("Field %s is required", field)
}

func (m mockMessages) Min(field string, min int) string {
	return fmt.Sprintf("The length of the field %s must be at least %d ", field, min)
}

func (m mockMessages) Max(field string, max int) string {
	return fmt.Sprintf("The length of the %s field must not be greater than %d", field, max)
}

func (m mockMessages) Email(field string) string {
	return fmt.Sprintf("The %s field must match the mail format", field)
}

func (m mockMessages) In(field string, in []string) string {
	inStr := strings.Join(in, " or ")
	return fmt.Sprintf("Field %s must have value %s", field, inStr)
}

func TestValidate(t *testing.T) {
	t.Run("test should be return err if struct does not contains fields", func(t *testing.T) {
		err := validator.Validate(mockUser{})

		if err == nil {
			t.Errorf("should return error")
		}
		t.Log(err)
	})

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

	t.Run("test should be return error if min or max values is not int type", func(t *testing.T) {
		err := validator.Validate(mockWrongUser{
			Email:    "example@example.com",
			Password: "p4ssw0rd",
		})

		if err == nil {
			t.Error("error should not be nil")
		}
		t.Log(err)
	})

	t.Run("test should be return nil if field passed max validation rules", func(t *testing.T) {
		err := validator.Validate(mockUser{
			Email:    "example@example.com",
			Password: "p4ssw0rd",
		})

		if err != nil {
			t.Error("error should be nil")
		}

		t.Log(err)
	})

	t.Run("test should be return error if 'in' rule not passed validation", func(t *testing.T) {
		err := validator.Validate(mockProduct{
			Name: "example",
			Type: "9",
		})

		if err == nil {
			t.Error("error should not be nil")
		}
		t.Log(err)
	})

	t.Run("test should be return nil if 'in' rule passed validation", func(t *testing.T) {
		err := validator.Validate(mockProduct{
			Name: "example",
			Type: "something",
		})

		if err != nil {
			t.Error("error should be nil")
		}
		t.Log(err)
	})
}
