package mappers

import (
	"net/http"
)

func Int8(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(request.PostForm.Get(hashName), 8, strict); err == nil {

		return int8(value), nil

	} else {

		return nil, err
	}
}
