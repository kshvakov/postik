package mappers

import (
	"github.com/kshvakov/errors"
	"net/http"
	"strconv"
)

func Float32(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if value, err := strconv.ParseFloat(request.PostForm.Get(hashName), 32); strict && err != nil {

		return 0, errors.Wrap(err)

	} else {

		return float32(value), nil
	}
}
