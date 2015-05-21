package mappers

import (
	"github.com/kshvakov/errors"
	"net/http"
	"strconv"
)

func SliceFloat32(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if values, err := values(hashName, request, strict); err == nil {

		result := make([]float32, 0, len(values))

		for _, val := range values {

			if v, err := strconv.ParseFloat(val, 32); err == nil {

				result = append(result, float32(v))

			} else {

				if strict {

					return nil, errors.Wrap(err)
				}
			}
		}

		return result, nil

	} else {

		return nil, err
	}
}
