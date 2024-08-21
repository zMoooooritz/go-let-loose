package util

import (
	"strconv"
)

func ToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return res
}

func ToInt64(str string) int64 {
	res, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return res
}
