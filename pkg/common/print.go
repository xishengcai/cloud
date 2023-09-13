package common

import (
	"encoding/json"

	uuid "github.com/satori/go.uuid"
)

func PrettifyJson(i interface{}, indent bool) string {
	var str []byte
	if indent {
		str, _ = json.MarshalIndent(i, "", "    ")
	} else {
		str, _ = json.Marshal(i)
	}

	return string(str)
}

func GetUUID() string {
	return uuid.NewV4().String()
}
