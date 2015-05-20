package mappers

import (
	"net/http"
)

func Int64(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(hashName, request, 64, strict); err == nil {

		return value, nil

	} else {

		return nil, err
	}
}
