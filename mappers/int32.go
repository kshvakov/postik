package mappers

import (
	"net/http"
)

func Int32(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(request.PostForm.Get(hashName), 32, strict); err == nil {

		return int32(value), nil

	} else {

		return nil, err
	}
}
