package bean

import (
	"github.com/json-iterator/go"
)

type RespBase struct {
	Msg string

	Code int
}

func ToString(req RespBase) string {
	b, err := jsoniter.Marshal(req)
	if err != nil {
		return ""
	}
	return string(b)
}
