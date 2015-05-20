package mappers

import (
	"net/http"
)

func Bool(hashName string, request *http.Request, strict bool) (interface{}, error) {

	value := request.PostForm.Get(hashName)

	return value != "" && value != "0" && value != "false", nil
}
