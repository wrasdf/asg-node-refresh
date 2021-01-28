package utils

import (
	"strconv"
	"encoding/json"
)

func StringToInt64(s string) (int64, error) {
  result, err := strconv.ParseInt(s, 10, 64)
  if err != nil {
    return 0, err
  }
  return result, nil
}

func ToJsonString(data interface{}) (string, error) {
	results, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(results), nil
}

func JsonStringToMap (s string) map[string]interface{} {
  var results map[string]interface{}
  json.Unmarshal([]byte(s), &results)
  return results
}
