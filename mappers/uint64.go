package mappers

import (
	"net/http"
)

func Uint64(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseUint(request.PostForm.Get(hashName), 64, strict); err == nil {

		return value, nil

	} else {

		return nil, err
	}
}
