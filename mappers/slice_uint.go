package mappers

import (
	"fmt"
	"github.com/kshvakov/errors"
	"net/http"
	"strconv"
)

func SliceUint(hashName string, request *http.Request, strict bool) (interface{}, error) {

	values := request.PostForm[fmt.Sprintf("%s[]", hashName)]

	result := make([]uint, 0, len(values))

	for _, val := range values {

		if v, err := strconv.ParseUint(val, 10, 0); err == nil {

			result = append(result, uint(v))
		} else {

			return nil, errors.Wrap(err)
		}
	}

	return result, nil
}
