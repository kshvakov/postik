package mappers

import (
	"fmt"
	"net/http"
)

func SliceString(hashName string, request *http.Request, strict bool) (interface{}, error) {

	return request.PostForm[fmt.Sprintf("%s[]", hashName)], nil
}
