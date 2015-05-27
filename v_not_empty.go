package postik

import (
	"fmt"
)

func NotEmpty(message string) Validator {
	return &notEmpty{
		message: message,
	}
}

type notEmpty struct {
	message string
}

func (_ *notEmpty) IsValid(field *Field) bool {

	v := fmt.Sprint(field.Value)

	return v != "" && v != "0"
}

func (l *notEmpty) Error() string {

	return l.message
}
