package mappers

import (
	"net/http"
)

func SliceString(hashName string, request *http.Request, strict bool) (interface{}, error) {

	return values(hashName, request, strict)
}
