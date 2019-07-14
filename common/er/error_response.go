package er

var (
	NoError = errorResponse{Status: true, Msg:"", Code: 0}
	ParamsError = errorResponse{Status: false, Msg:"传入参数不对", Code: 1}
	JsonBodyError = errorResponse{Status: false, Msg:"传入body格式不对", Code: 2}
	SourceNotExistError = errorResponse{Status: false, Msg:"资源不存在", Code: 404}

	UserNotExistError = errorResponse{Status: false, Msg:"用户不存在", Code: 1000}
	PasswordError = errorResponse{Status: false, Msg:"用户密码不对", Code: 1001}
)

type errorResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
	Msg string `json:"msg"`
	Code int `json:"code"`
}

func (error errorResponse) SetData(data interface{}) errorResponse{
	error.Data = data
	return error
}
