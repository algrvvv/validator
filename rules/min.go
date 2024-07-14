package rules

import (
	"errors"
	"fmt"
	"github.com/algrvvv/validator/types"
)

// Min function to check the minimum field length
func Min(f, tag, v string, m types.IMessages) error {
	var minValue int
	_, err := fmt.Sscanf(tag, "min=%d", &minValue)
	if err != nil {
		// паникуем, чтобы недопустить использование
		// неправильного минимального значения для работы
		panic("ошибка чтения минимального значения")
	}

	if minValue > len(v) {
		return errors.New(m.Min(f, minValue))
	}

	return nil
}
