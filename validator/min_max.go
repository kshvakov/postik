package validator

import (
	"github.com/kshvakov/postik"
)

func MinMax(message string, min int, max uint) postik.Validator {
	return &minMax{
		message: message,
		min:     min,
		max:     max,
	}
}

type minMax struct {
	message string
	min     int
	max     uint
}

func (m *minMax) IsValid(field *postik.Field) bool {

	switch field.Value.(type) {

	case int:

		v := field.Value.(int)

		return v >= m.min && v <= int(m.max)

	case int8:

		v := int(field.Value.(int8))

		return v >= m.min && v <= int(m.max)

	case int16:

		v := int(field.Value.(int16))

		return v >= m.min && v <= int(m.max)

	case int32:

		v := int(field.Value.(int32))

		return v >= m.min && v <= int(m.max)

	case int64:

		v := field.Value.(int64)

		return v >= int64(m.min) && v <= int64(m.max)

	case uint:

		v := field.Value.(uint)

		return v >= uint(m.min) && v <= m.max

	case uint8:

		v := uint(field.Value.(uint8))

		return v >= uint(m.min) && v <= m.max

	case uint16:

		v := uint(field.Value.(uint16))

		return v >= uint(m.min) && v <= m.max

	case uint32:

		v := uint(field.Value.(uint32))

		return v >= uint(m.min) && v <= m.max

	case uint64:

		v := field.Value.(uint64)

		return v >= uint64(m.min) && v <= uint64(m.max)
	}

	return false
}

func (m *minMax) Error() string {

	return m.message
}
