package mappers

import (
	"net/http"
)

func Uint(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(request.PostForm.Get(hashName), 0, strict); err == nil {

		return uint(value), nil

	} else {

		return nil, err
	}
}
