package mappers

import (
	"net/http"
)

func Int16(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(request.PostForm.Get(hashName), 16, strict); err == nil {

		return int16(value), nil

	} else {

		return nil, err
	}
}
