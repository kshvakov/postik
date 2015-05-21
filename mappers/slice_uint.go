package mappers

import (
	"net/http"
)

func SliceUint(hashName string, request *http.Request, strict bool) (interface{}, error) {

	if values, err := values(hashName, request, strict); err == nil {

		result := make([]uint, 0, len(values))

		for _, val := range values {

			if v, err := parseUint(val, 10, strict); err == nil {

				result = append(result, uint(v))

			} else {

				return nil, err
			}
		}

		return result, nil

	} else {

		return nil, err
	}
}
