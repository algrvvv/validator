package types

import "fmt"

type IMessages interface {
	Required(field string) string
	Min(field string, min int) string
	Email(field string) string
}

type DefaultMessages struct {
}


func (d DefaultMessages) Required(field string) string {
	return fmt.Sprintf("Поле %s обязательно", field)
}

func (d DefaultMessages) Min(field string, min int) string {
	return fmt.Sprintf("Длина поля %s должна быть не меньше %d ", field, min)
}

func (d DefaultMessages) Email(field string) string {
	return fmt.Sprintf("Поле %s должно соответствовать формату почты", field)
}
