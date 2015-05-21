package mappers

import (
	"net/http"
)

func Uint16(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(request.PostForm.Get(hashName), 16, strict); err == nil {

		return uint16(value), nil

	} else {

		return nil, err
	}
}
