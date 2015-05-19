package mappers

import (
	"github.com/kshvakov/errors"
	"net/http"
	"strconv"
)

func Int(hashName string, request *http.Request) (interface{}, error) {

	if value, err := strconv.ParseInt(request.PostForm.Get(hashName), 10, 0); err == nil {

		return int(value), nil

	} else {

		return 0, errors.Wrap(err)
	}
}
