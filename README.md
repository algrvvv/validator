# Validator - p4xt3r

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
