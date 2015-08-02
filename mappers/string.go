package mappers

import (
	"fmt"
	"net/http"
)

func String(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if _, found := request.PostForm[hashName]; !found && strict {

		return nil, fmt.Errorf("key '%s' not found", hashName)
	}

	return request.PostForm.Get(hashName), nil
}
