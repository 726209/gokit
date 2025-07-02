package common

import (
	"encoding/json"
	"strings"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func PrettyJSONString(data interface{}) string {
	str, _ := PrettyStruct(data)
	return str
}

func JSONString(data interface{}) string {
	str, _ := PrettyStruct(data)
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

// Capitalize 首字母大写
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
