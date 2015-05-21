package mappers

import (
	"net/http"
)

func Int(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(request.PostForm.Get(hashName), 0, strict); err == nil {

		return int(value), nil

	} else {

		return nil, err
	}
}
