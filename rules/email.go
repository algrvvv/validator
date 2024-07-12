package rules

import (
	"errors"
	"regexp"
	
	"github.com/algrvvv/validator/types"
)

func Email(e,v string, m types.IMessages) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)

	if re.MatchString(v) {
		return nil
	} else {
		return errors.New(m.Email(e))
	}
}
