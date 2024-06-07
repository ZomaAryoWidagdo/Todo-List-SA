package utils

import (
	"strconv"
)

func StringToUint(str string) (uint, error) {
	num64, err := strconv.ParseUint(str, 10, 64)

	if err != nil {
		return 0, err
	}

	num := uint(num64)

	return num, nil
}