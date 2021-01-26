package utils

import (
	"strconv"
)

func StringToInt64(s string) (int64, error) {
  result, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return result, nil
	}
  return 0, err
}
