package re

import (
	"stouch_server/src/common/er"
)

type ResponseResult struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Code   int         `json:"code"`
}

func Data(data interface{}) ResponseResult {
	return ResponseResult{Status: true, Msg: "", Data: data, Code: 0}
}

func Error(error er.Error) ResponseResult {
	return ResponseResult{Status: false, Msg: error.Msg, Code: error.Code}
}
