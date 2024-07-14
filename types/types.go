package types

import "fmt"

type IMessages interface {
	// Required method with a message when there is an error checking whether a field is required
	Required(field string) string

	// Min method that sends a message when there is an error in validating the minimum field length
	Min(field string, min int) string

	// Email method that sends a message when there is a field validation error with the mail format
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
