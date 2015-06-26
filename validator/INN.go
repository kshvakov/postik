package validator

import (
	"fmt"
	"github.com/kshvakov/inn"
	"github.com/kshvakov/postik"
)

func INN(message string) postik.Validator {
	return &checkINN{
		message: message,
	}
}

type checkINN struct {
	message string
}

func (c *checkINN) IsValid(field *postik.Field) bool {

	return inn.Check(fmt.Sprint(field.Value))
}

func (c *checkINN) Error() string {

	return c.message
}
