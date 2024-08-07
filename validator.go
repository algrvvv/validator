package validator

import (
	"reflect"
	"strings"

	"github.com/algrvvv/validator/rules"
	"github.com/algrvvv/validator/types"
)

// Validate function for structure validation
// uses default error messages
func Validate(s interface{}) error {
	return ValidateWithMessage(s, &types.DefaultMessages{})
}

// ValidateWithMessage function for structure validation with messages
// uses IMessages implementation for show custom error messages
func ValidateWithMessage(s interface{}, m types.IMessages) error {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tags := field.Tag.Get("validate")

		if tags == "" {
			return nil
		}

		tagList := strings.Split(tags, ",")
		for _, tag := range tagList {
			switch {
			case tag == "required":
				if err := rules.Required(field.Name, value.Interface(), m); err != nil {
					return err
				}
			case tag == "email":
				if err := rules.Email(field.Name, value.String(), m); err != nil {
					return err
				}
			case strings.HasPrefix(tag, "min="):
				if err := rules.Min(field.Name, tag, value.String(), m); err != nil {
					return err
				}
			case strings.HasPrefix(tag, "max="):
				if err := rules.Max(field.Name, tag, value.String(), m); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
