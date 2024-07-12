package rules

import (
	"errors"

	"github.com/algrvvv/validator/types"
)

func Required(f string, v any, m types.IMessages) error {
	if v == "" {
		return errors.New(m.Required(f))
	}

	return nil
}
