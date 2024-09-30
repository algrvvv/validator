# Validator - algrvvv

## Установка

```shell
$ go get github.com/algrvvv/validator
```

```go
import "github.com/algrvvv/validator"
```

## Использование

Список полей поддерживающихся на данный момент:
- required - обязательное поле
- email - поле должно быть формата email
- min=x - минимальная длина, где x длина
- max=x - максимальная длина, где x длина
- in:x-y-z - сравнение значения поля с x,y,z

```go
// создаем струтуру с нужными полями
type mockUser struct {
    Email    string `validate:"required,email"`
    Password string `validate:"required,min=8"`
}

// nil, если ошибок валидации нет
// err, если есть ошибка валидации.
// Err будет содержать в себе текст ошибки
err := validator.Validate(mockUser{
    Email:    "example@example.com",
    Password: "p4ssw0rd",
})

```

### Использование своих сообщений об ошибке

```go
// Требуется создать структуру, которая имплементирует IMessage
type IMessages interface {
    Required(field string) string
    Min(field string, min int) string
    Email(field string) string
}

// где field - поле, которое не прошло проверку
// А использовать вот так:
err := validator.ValidateWithMessage(mockUser{
    Email:    "example@example.com",
    Password: "p4ssw0rd",
}, &mockMessages{})
```

Более подробный пример можно увидеть в `validator_test.go`
