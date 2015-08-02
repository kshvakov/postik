package mappers

import (
	"fmt"
	"net/http"
	"strconv"
)

func parseInt(value string, bitSize int, strict bool) (int64, error) {

	if value, err := strconv.ParseInt(value, 10, bitSize); strict && err != nil {

		return 0, err

	} else {

		return value, nil
	}
}

func parseUint(value string, bitSize int, strict bool) (uint64, error) {

	if value, err := strconv.ParseUint(value, 10, bitSize); strict && err != nil {

		return 0, err

	} else {

		return value, nil
	}
}

func values(hashName string, request *http.Request, strict bool) ([]string, error) {

	if values, found := request.PostForm[fmt.Sprintf("%s[]", hashName)]; found {

		return values, nil

	} else {

		if strict {

			return nil, fmt.Errorf("key '%s' not found", hashName)
		}

		return []string{}, nil
	}
}
