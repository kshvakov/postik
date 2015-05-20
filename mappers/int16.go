package mappers

import (
	"net/http"
)

func Int16(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := parseInt(hashName, request, 16, strict); err == nil {

		return int(value), nil

	} else {

		return nil, err
	}
}
