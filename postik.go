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

func New(obj interface{}, params ...string) *postik {

	var (
		t = DefaultTag
		s = ""
	)

	switch len(params) {
	case 1:

		s = params[0]

	case 2:

		s = params[0]
		t = params[1]

	}

	p := &postik{
		tag:    t,
		salt:   s,
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

		if name, strict := tag(fl.Tag.Get(p.tag)); name != "" && name != "-" {

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
	tag          string
	salt         string
	excludeNames []string
	fields       map[string]*Field
}

func (p *postik) SetValidators(validators map[string][]Validator) {

	for name, v := range validators {

		if field, found := p.fields[name]; found {

			empty := reflect.TypeOf(&notEmpty{})

			for _, r := range v {

				if empty == reflect.TypeOf(r) {

					field.Required = true

					break
				}
			}

			field.validators = v
		}
	}
}

func (p *postik) ExcludeNames(names ...string) {

	p.excludeNames = names
}

func (p *postik) Fields() map[string]*Field {

	return p.fields
}

func (p *postik) exclude(name string) bool {

	for _, n := range p.excludeNames {

		if n == name {

			return true
		}
	}

	return false
}

func (p *postik) ParseForm(request *http.Request) error {

	if err := request.ParseForm(); err != nil {

		return err
	}

	for name, field := range p.fields {

		if !p.exclude(name) {

			if err := field.ParseForm(request); err != nil {

				return err
			}
		}
	}

	return nil
}

func (p *postik) IsValid() bool {

	isValid := true

	for name, field := range p.fields {

		if !p.exclude(name) && !field.IsValid() {

			isValid = false
		}
	}

	if isValid {

		for name, field := range p.fields {

			if !p.exclude(name) {

				field.parent.Set(reflect.ValueOf(field.Value))
			}
		}
	}

	return isValid
}

func tag(tag string) (string, bool) {

	if tags := strings.Split(tag, ","); len(tags) > 0 {

		for _, t := range tags {

			if t == "strict" {

				return tags[0], true
			}
		}

		return tags[0], false
	}

	return "", false
}
