package error_response

var NoError = errorResponse{Status: true, Msg:"", Code: 0}

type errorResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}
