package mappers

import (
	"net/http"
)

func Uint32(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(request.PostForm.Get(hashName), 32, strict); err == nil {

		return uint32(value), nil

	} else {

		return nil, err
	}
}
