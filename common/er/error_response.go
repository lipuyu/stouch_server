package er

var (
	NoError = errorResponse{Status: true, Msg:"", Code: 0}
	ParamsError = errorResponse{Status: false, Msg:"传入参数不对", Code: 1}
	JsonBodyError = errorResponse{Status: false, Msg:"传入body格式不对", Code: 2}
	AppError = errorResponse{Status: false, Msg:"App名字不对", Code: 3}

	SourceNotExistError = errorResponse{Status: false, Msg:"资源不存在", Code: 404}

	UserNotExistError = errorResponse{Status: false, Msg:"用户不存在", Code: 1000}
	PasswordError = errorResponse{Status: false, Msg:"用户密码不对", Code: 1001}
	UserNameRepeatError = errorResponse{Status: false, Msg:"用户名已经存在", Code: 1002}
)

type errorResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func Data(data interface{}) errorResponse {
	return errorResponse{Status: true, Msg:"", Data:data, Code: 0}
}