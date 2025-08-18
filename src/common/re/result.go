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

func Error(err er.Error) ResponseResult {
	return ResponseResult{Status: false, Msg: err.Msg, Code: err.Code}
}

func ErrorStr(err error) ResponseResult {
	return ResponseResult{Status: false, Msg: err.Error(), Code: -1}
}
