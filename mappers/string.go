package mappers

import (
	"github.com/kshvakov/errors"
	"net/http"
)

func String(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if _, found := request.PostForm[hashName]; !found && strict {

		return nil, errors.New("value not found")
	}

	return request.PostForm.Get(hashName), nil
}
