package rules

import (
	"errors"
	"fmt"

	"github.com/algrvvv/validator/types"
)

func Max(field, tag, value string, m types.IMessages) error {
	var (
		maxValue int
		err      error
	)

	_, err = fmt.Sscanf(tag, "max=%d", &maxValue)
	if err != nil {
		return errors.New("ошибка чтения максимального значения")
	}

	if maxValue < len(value) {
		return errors.New(m.Max(field, maxValue))
	}

	return nil
}
