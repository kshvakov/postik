package mappers

import (
	"net/http"
)

func Uint8(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(hashName, request, 8, strict); err == nil {

		return uint8(value), nil

	} else {

		return nil, err
	}
}
