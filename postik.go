package postik

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

const (
	DefaultTag = "form"
)

func New(obj interface{}, salt ...string) *postik {

	p := &postik{
		Tag:    DefaultTag,
		salt:   strings.Join(salt, "."),
		fields: make(map[string]*Field),
	}

	tp := reflect.TypeOf(obj)

	if tp.Kind() == reflect.Ptr {

		tp = tp.Elem()
	}

	el := reflect.ValueOf(obj)

	if el.Kind() == reflect.Ptr {

		el = el.Elem()
	}

	for i := 0; i < tp.NumField(); i++ {

		fl := tp.Field(i)

		if fl.PkgPath != "" {

			continue
		}

		fd := el.Field(i)

		if fd.Kind() == reflect.Ptr {

			fd = fd.Elem()
		}

		if name, strict := tag(fl.Tag.Get(p.Tag)); name != "" && name != "-" {

			p.fields[name] = &Field{
				Name:     name,
				HashName: hash(fmt.Sprintf("%s%s", name, p.salt)),
				Value:    fd.Interface(),
				parent:   fd,
				mapper:   Map(fd.Interface()),
				strict:   strict,
			}
		}
	}

	return p
}

type postik struct {
	Tag    string
	salt   string
	fields map[string]*Field
}

func (p *postik) SetValidators(validators map[string][]Validator) {

	for name, v := range validators {

		if field, found := p.fields[name]; found {

			field.validators = v
		}
	}
}

func (p *postik) Fields() map[string]*Field {

	return p.fields
}

func (p *postik) ParseForm(request *http.Request) error {

	if err := request.ParseForm(); err != nil {

		return err
	}

	for _, field := range p.fields {

		if err := field.ParseForm(request); err != nil {

			return err
		}
	}

	return nil
}

func (p *postik) IsValid() bool {

	isValid := true

	for _, field := range p.fields {

		if !field.IsValid() {

			isValid = false
		}
	}

	if isValid {

		for _, field := range p.fields {

			field.parent.Set(reflect.ValueOf(field.Value))
		}
	}

	return isValid
}

func tag(tag string) (string, bool) {

	if tags := strings.Split(tag, ","); len(tags) > 0 {

		var strict bool

		for _, t := range tags {

			if t == "strict" {

				strict = true

				break
			}
		}

		return tags[0], strict
	}

	return "", false
}
