package mappers

import (
	"net/http"
)

func String(hashName string, request *http.Request) (interface{}, error) {

	return request.PostForm.Get(hashName), nil
}
