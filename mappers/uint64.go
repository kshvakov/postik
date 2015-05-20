package mappers

import (
	"net/http"
)

func Uint64(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(hashName, request, 64, strict); err == nil {

		return value, nil

	} else {

		return nil, err
	}
}
