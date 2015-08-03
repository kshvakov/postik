package validator

import (
	"github.com/kshvakov/postik"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMathMinMaxInt8(t *testing.T) {

	validator := MinMax("", math.MinInt8+2, math.MaxInt8-1)

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int8(math.MinInt8 + 3),
	}))

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int8(math.MaxInt8 - 2),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int8(math.MinInt8),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int8(math.MaxInt8),
	}))
}

func TestMathMinMaxInt16(t *testing.T) {

	validator := MinMax("", math.MinInt16+2, math.MaxInt16-1)

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int16(math.MinInt16 + 3),
	}))

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int16(math.MaxInt16 - 2),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int16(math.MinInt16),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int16(math.MaxInt16),
	}))
}

func TestMathMinMaxInt32(t *testing.T) {

	validator := MinMax("", math.MinInt32+2, math.MaxInt32-1)

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int32(math.MinInt32 + 3),
	}))

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int32(math.MaxInt32 - 2),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int32(math.MinInt32),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int32(math.MaxInt32),
	}))
}

func TestMathMinMaxInt64(t *testing.T) {

	validator := MinMax("", math.MinInt64+2, math.MaxInt64-1)

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int64(math.MinInt64 + 3),
	}))

	assert.True(t, validator.IsValid(&postik.Field{
		Value: int64(math.MaxInt64 - 2),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int64(math.MinInt64),
	}))

	assert.False(t, validator.IsValid(&postik.Field{
		Value: int64(math.MaxInt64),
	}))
}

func TestMinMaxInt(t *testing.T) {

	validator := MinMax("", -42, 42)

	valid := []interface{}{
		int(42),
		int8(42),
		int16(42),
		int32(42),
		int64(42),
		int(21),
		int8(21),
		int16(21),
		int32(21),
		int64(21),

		int(-42),
		int8(-42),
		int16(-42),
		int32(-42),
		int64(-42),
		int(-21),
		int8(-21),
		int16(-21),
		int32(-21),
		int64(-21),
	}

	for _, v := range valid {

		assert.True(t, validator.IsValid(&postik.Field{
			Value: v,
		}))
	}

	invalid := []interface{}{
		int(43),
		int8(43),
		int16(43),
		int32(43),
		int64(43),

		int(-43),
		int8(-43),
		int16(-43),
		int32(-43),
		int64(-43),
	}

	for _, v := range invalid {

		assert.False(t, validator.IsValid(&postik.Field{
			Value: v,
		}))
	}
}

func TestMinMaxUint(t *testing.T) {

	validator := MinMax("", 10, 42)

	valid := []interface{}{
		uint(42),
		uint8(42),
		uint16(42),
		uint32(42),
		uint64(42),
		uint(21),
		uint8(21),
		uint16(21),
		uint32(21),
		uint64(21),

		uint(10),
		uint8(10),
		uint16(10),
		uint32(10),
		uint64(10),
		uint(20),
		uint8(20),
		uint16(20),
		uint32(20),
		uint64(20),
	}

	for _, v := range valid {

		assert.True(t, validator.IsValid(&postik.Field{
			Value: v,
		}))
	}

	invalid := []interface{}{
		uint(43),
		uint8(43),
		uint16(43),
		uint32(43),
		uint64(43),

		uint(5),
		uint8(5),
		uint16(5),
		uint32(5),
		uint64(5),
		uint(9),
		uint8(9),
		uint16(9),
		uint32(9),
		uint64(9),
	}

	for _, v := range invalid {

		assert.False(t, validator.IsValid(&postik.Field{
			Value: v,
		}))
	}
}

func BenchmarkMinMax(b *testing.B) {

	b.ReportAllocs()

	validator := MinMax("", 10, 42)

	values := []interface{}{
		int(42),
		int8(42),
		int16(42),
		int32(42),
		int64(42),
		int(21),
		int8(21),
		int16(21),
		int32(21),
		int64(21),

		int(-42),
		int8(-42),
		int16(-42),
		int32(-42),
		int64(-42),
		int(-21),
		int8(-21),
		int16(-21),
		int32(-21),
		int64(-21),

		uint(42),
		uint8(42),
		uint16(42),
		uint32(42),
		uint64(42),
		uint(21),
		uint8(21),
		uint16(21),
		uint32(21),
		uint64(21),

		uint(10),
		uint8(10),
		uint16(10),
		uint32(10),
		uint64(10),
		uint(20),
		uint8(20),
		uint16(20),
		uint32(20),
		uint64(20),
	}

	b.ResetTimer()

	var v int

	for i := 0; i < b.N; i++ {

		_ = validator.IsValid(&postik.Field{
			Value: values[v],
		})

		v++

		if v >= len(values)-1 {
			v = 0
		}
	}
}
