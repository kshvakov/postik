package mappers

import (
	"github.com/kshvakov/errors"
	"net/http"
	"strconv"
)

func parseInt(hashName string, request *http.Request, bitSize int, strict bool) (int64, error) {

	if value, err := strconv.ParseInt(request.PostForm.Get(hashName), 10, bitSize); strict && err != nil {

		return 0, errors.Wrap(err)

	} else {

		return value, nil
	}
}

func parseUint(hashName string, request *http.Request, bitSize int, strict bool) (uint64, error) {

	if value, err := strconv.ParseUint(request.PostForm.Get(hashName), 10, bitSize); strict && err != nil {

		return 0, errors.Wrap(err)

	} else {

		return value, nil
	}
}
