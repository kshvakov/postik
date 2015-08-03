package postik

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"net/http"
	"net/url"
	"testing"
)

type builtin struct {
	Bool bool `form:"boolean"`

	String string `form:"string"`
	//SliceString []string `form:"slice_string"`

	MinInt8  int8  `form:"min_int8"`
	MaxInt8  int8  `form:"max_int8"`
	MinInt16 int16 `form:"min_int16"`
	MaxInt16 int16 `form:"max_int16"`
	MinInt32 int32 `form:"min_int32"`
	MaxInt32 int32 `form:"max_int32"`
	MinInt64 int64 `form:"min_int64"`
	MaxInt64 int64 `form:"max_int64"`
	MinInt   int   `form:"min_int"`
	MaxInt   int   `form:"max_int"`

	MaxUint8  uint8  `form:"uint8"`
	MaxUint16 uint16 `form:"uint16"`
	MaxUint32 uint32 `form:"uint32"`
	MaxUint64 uint64 `form:"uint64"`
	MaxUint   uint   `form:"uint"`

	MinFloat32 float32 `form:"min_float32"`
	MaxFloat32 float32 `form:"max_float32"`
	MinFloat64 float64 `form:"min_float64"`
	MaxFloat64 float64 `form:"max_float64"`
}

func TestBuiltinTypes(t *testing.T) {

	buffer := new(bytes.Buffer)

	params := url.Values{}

	params.Set("boolean", "true")

	params.Set("string", "String")
	//params["slice_string[]"] = []string{"s1", "s2", "s3", "s4", "s5"}

	params.Set("min_int8", fmt.Sprint(math.MinInt8))
	params.Set("max_int8", fmt.Sprint(math.MaxInt8))
	params.Set("min_int16", fmt.Sprint(math.MinInt16))
	params.Set("max_int16", fmt.Sprint(math.MaxInt16))
	params.Set("min_int32", fmt.Sprint(math.MinInt32))
	params.Set("max_int32", fmt.Sprint(math.MaxInt32))
	params.Set("min_int64", fmt.Sprint(math.MinInt64))
	params.Set("max_int64", fmt.Sprint(math.MaxInt64))
	params.Set("min_int", fmt.Sprint(math.MinInt64))
	params.Set("max_int", fmt.Sprint(math.MaxInt64))

	params.Set("uint8", fmt.Sprint(math.MaxUint8))
	params.Set("uint16", fmt.Sprint(math.MaxUint16))
	params.Set("uint32", fmt.Sprint(math.MaxUint32))
	params.Set("uint64", fmt.Sprint(uint64(math.MaxUint64)))
	params.Set("uint", fmt.Sprint(uint64(math.MaxUint64)))

	params.Set("min_float32", fmt.Sprint(float32(math.SmallestNonzeroFloat32)))
	params.Set("max_float32", fmt.Sprint(float32(math.MaxFloat32)))
	params.Set("min_float64", fmt.Sprint(float64(math.SmallestNonzeroFloat64)))
	params.Set("max_float64", fmt.Sprint(float64(math.MaxFloat64)))

	buffer.WriteString(params.Encode())

	req, _ := http.NewRequest("POST", "", buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	data := builtin{}

	form := New(&data)

	err := form.ParseForm(req)

	if assert.NoError(t, err) && assert.True(t, form.IsValid()) {

		assert.True(t, data.Bool)

		assert.Equal(t, "String", data.String)
		//assert.Equal(t, []string{"s1", "s2", "s3", "s4", "s5"}, data.SliceString)

		assert.Equal(t, math.MinInt8, int(data.MinInt8))
		assert.Equal(t, math.MaxInt8, int(data.MaxInt8))
		assert.Equal(t, math.MinInt16, int(data.MinInt16))
		assert.Equal(t, math.MaxInt16, int(data.MaxInt16))
		assert.Equal(t, math.MinInt32, int(data.MinInt32))
		assert.Equal(t, math.MaxInt32, int(data.MaxInt32))
		assert.Equal(t, math.MinInt64, int(data.MinInt64))
		assert.Equal(t, math.MaxInt64, int(data.MaxInt64))
		assert.Equal(t, math.MinInt64, int(data.MinInt))
		assert.Equal(t, math.MaxInt64, int(data.MaxInt))

		assert.Equal(t, uint(math.MaxUint8), uint(data.MaxUint8))
		assert.Equal(t, uint(math.MaxUint16), uint(data.MaxUint16))
		assert.Equal(t, uint(math.MaxUint32), uint(data.MaxUint32))
		assert.Equal(t, uint64(math.MaxUint64), data.MaxUint64)
		assert.Equal(t, uint(math.MaxUint64), uint(data.MaxUint))

		assert.Equal(t, float32(math.SmallestNonzeroFloat32), data.MinFloat32)
		assert.Equal(t, float32(math.MaxFloat32), data.MaxFloat32)
		assert.Equal(t, float64(math.SmallestNonzeroFloat64), data.MinFloat64)
		assert.Equal(t, float64(math.MaxFloat64), data.MaxFloat64)

		//fmt.Printf("%#v\n", data)

	}
}
