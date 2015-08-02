package postik

import (
	mapper "github.com/kshvakov/postik/mappers"
	"net/http"
	"reflect"
)

var mappers map[reflect.Type]Mapper

type Mapper func(hashName string, request *http.Request, strict bool) (interface{}, error)

func Map(t interface{}) Mapper {

	if mapper, found := mappers[reflect.TypeOf(t)]; found {

		return mapper
	}

	panic(reflect.TypeOf(t))
}

func init() {

	mappers = map[reflect.Type]Mapper{
		reflect.TypeOf(int(1)):       mapper.Int,
		reflect.TypeOf(int8(1)):      mapper.Int8,
		reflect.TypeOf(int16(1)):     mapper.Int16,
		reflect.TypeOf(int32(1)):     mapper.Int32,
		reflect.TypeOf(int64(1)):     mapper.Int64,
		reflect.TypeOf(float32(1)):   mapper.Float32,
		reflect.TypeOf(float64(1)):   mapper.Float64,
		reflect.TypeOf(uint(1)):      mapper.Uint,
		reflect.TypeOf(uint8(1)):     mapper.Uint8,
		reflect.TypeOf(uint16(1)):    mapper.Uint16,
		reflect.TypeOf(uint32(1)):    mapper.Uint32,
		reflect.TypeOf(uint64(1)):    mapper.Uint64,
		reflect.TypeOf(""):           mapper.String,
		reflect.TypeOf([]string{""}): mapper.SliceString,
		reflect.TypeOf([]uint{1}):    mapper.SliceUint,
		reflect.TypeOf([]float32{1}): mapper.SliceFloat32,
		reflect.TypeOf(true):         mapper.Bool,
	}
}

func RegisterMapper(t interface{}, mapper Mapper) {

	mappers[reflect.TypeOf(t)] = mapper
}
