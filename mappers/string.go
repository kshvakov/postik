package mappers

import (
	"net/http"
)

func String(hashName string, request *http.Request, strict bool) (interface{}, error) {

	return request.PostForm.Get(hashName), nil
}
