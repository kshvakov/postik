package validator

import (
	"fmt"
	"github.com/kshvakov/postik"
)

func Len(message string, min, max int) postik.Validator {
	return &lenMinMax{
		message: message,
		min:     min,
		max:     max,
	}
}

type lenMinMax struct {
	message  string
	min, max int
}

func (l *lenMinMax) IsValid(field *postik.Field) bool {

	v := fmt.Sprint(field.Value)

	return len(v) >= l.min && len(v) <= l.max
}

func (l *lenMinMax) Error() string {

	return l.message
}
