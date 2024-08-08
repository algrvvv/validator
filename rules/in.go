package rules

import (
	"errors"
	"strings"

	"github.com/algrvvv/validator/types"
)

func In(field, tag, value string, m types.IMessages) error {
	inValues := strings.TrimPrefix(tag, "in:")

	splited := strings.Split(inValues, "-")

	for _, v := range splited {
		if v == value {
			return nil
		}
	}

	return errors.New(m.In(field, splited))
}
