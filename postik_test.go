package postik

import (
	"bytes"
	"fmt"
	"github.com/kshvakov/errors"
	"net/http"
	"net/url"
	"testing"
)

func Len(message string, min, max int) Validator {

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

func (l *lenMinMax) IsValid(field *Field) bool {

	v := fmt.Sprintf("%v", field.Value)

	return len(v) >= l.min && len(v) <= l.max
}

func (l *lenMinMax) Error() string {

	return l.message
}

type S struct {
	Name string `form:"name"`
	//NameStrict   string    `form:"name_strict,strict"`
	Value        int       `form:"value,strict"`
	CanBe        bool      `form:"can_be"`
	SliceUint    []uint    `form:"slice_uint"`
	SliceFloat32 []float32 `form:"slice_float32"`
	SliceString  []string  `form:"slice_string"`
	NoForm       []uint
}

func TestDev(t *testing.T) {

	s := S{Name: "AA", Value: 42}

	buffer := new(bytes.Buffer)

	params := url.Values{}

	params.Set("name", "Test322")
	params.Set("value", "8484")
	params.Set("can_be", "true")

	params["slice_uint[]"] = []string{"1", "2", "3", "4", "5"}
	params["slice_float32[]"] = []string{"1.1", "2.2", "3.23", "4.4", "5.5"}
	params["slice_string[]"] = []string{"s1", "s2", "s3", "s4", "s5"}

	buffer.WriteString(params.Encode())

	req, _ := http.NewRequest("POST", "", buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	p := New(&s)
	p.SetValidators(map[string][]Validator{
		"name": {
			Len("message", 5, 7),
		},
	})

	fields := p.Fields()

	fmt.Printf("%#v\n", fields)

	fmt.Printf("%v\n", s)

	if err := p.ParseForm(req); err == nil {

		if p.IsValid() {

			fmt.Printf("%v\n", s)

		} else {

			fmt.Println(req)

			for _, field := range fields {

				fmt.Printf("%s err: %s\n", field.Name, field.Error)
			}
		}

	} else {

		if e, ok := err.(errors.Error); ok {

			fmt.Println(e.Error(), e.Stack())

		} else {

			fmt.Println(err)
		}
	}
}

func BenchmarkDev(b *testing.B) {

	s := S{Name: "AA", Value: 42}

	buffer := new(bytes.Buffer)

	params := url.Values{}

	params.Set("name", "Test")
	params.Set("value", "8484")
	//params.Set("can_be", "false")
	params["slice_uint[]"] = []string{"1", "2", "3", "4", "5"}
	buffer.WriteString(params.Encode())

	req, _ := http.NewRequest("POST", "", buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	b.ReportAllocs()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		p := New(&s)
		p.ParseForm(req)
		p.IsValid()
	}
}
