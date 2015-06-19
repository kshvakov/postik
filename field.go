package postik

import (
	"fmt"
	"net/http"
	"reflect"
)

type Field struct {
	Name       string
	HashName   string
	Value      interface{}
	Error      string
	Required   bool
	validators []Validator
	parent     reflect.Value
	mapper     Mapper
	strict     bool
}

func (f *Field) ParseForm(request *http.Request) error {

	if value, err := f.mapper(f.HashName, request, f.strict); err == nil {

		f.Value = value

		return nil

	} else {

		return err
	}
}

func (f *Field) In(v interface{}) bool {

	if reflect.TypeOf(f.Value).Kind() == reflect.Slice {

		slice := reflect.ValueOf(f.Value)

		for i := 0; i < slice.Len(); i++ {

			if fmt.Sprint(slice.Index(i).Interface()) == fmt.Sprint(v) {

				return true
			}
		}
	}

	return false
}

func (f *Field) IsValid() bool {

	for _, validator := range f.validators {

		if !validator.IsValid(f) {

			f.Error = validator.Error()

			return false
		}
	}

	return true
}
