package hello

import (
	"github.com/tidwall/gjson"
)

func GetHello() string {
	json := `{"list":{"hello":"Hello","world":"World"}}`
	value := gjson.Get(json, "list.world")
	return value.String()
}
