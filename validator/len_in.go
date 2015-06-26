package validator

import (
	"fmt"
	"github.com/kshvakov/postik"
	"strings"
	"unicode/utf8"
)

func LenIn(message string, ln ...int) postik.Validator {

	return &lenIn{
		message: message,
		ln:      ln,
	}
}

type lenIn struct {
	message string
	ln      []int
}

func (l *lenIn) IsValid(field *postik.Field) bool {

	ln := utf8.RuneCountInString(fmt.Sprint(field.Value))

	return in(ln, l.ln)
}

func (l *lenIn) Error() string {

	var vv []interface{}

	for _, v := range l.ln {

		vv = append(vv, v)
	}

	if strings.Index(l.message, "%") != -1 {

		return fmt.Sprintf(l.message, vv...)
	}

	return l.message
}

func in(a int, b []int) bool {

	for _, z := range b {

		if z == a {

			return true
		}
	}

	return false
}
