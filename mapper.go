package postik

import (
	mapper "github.com/kshvakov/postik/mappers"
	"net/http"
	"reflect"
)

var mappers map[reflect.Type]Mapper

type Mapper func(hashName string, request *http.Request) (interface{}, error)

func Map(t interface{}) Mapper {

	if mapper, found := mappers[reflect.TypeOf(t)]; found {

		return mapper
	}

	panic(reflect.TypeOf(t))
}

func init() {

	mappers = map[reflect.Type]Mapper{
		reflect.TypeOf(int(1)):       mapper.Int,
		reflect.TypeOf(""):           mapper.String,
		reflect.TypeOf([]string{""}): mapper.SliceString,
		reflect.TypeOf([]uint{1}):    mapper.SliceUint,
		reflect.TypeOf(true):         mapper.Bool,
	}
}

func RegisterMapper(t interface{}, mapper Mapper) {

	mappers[reflect.TypeOf(t)] = mapper
}
