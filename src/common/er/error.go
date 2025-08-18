package er

var (
	ParamsError   = Error{Msg: "传入参数不对", Code: 1}
	JsonBodyError = Error{Msg: "传入body格式不对", Code: 2}
	AppError      = Error{Msg: "App名字不对", Code: 3}

	SourceNotExistError = Error{Msg: "资源不存在", Code: 404}

	UserNotExistError   = Error{Msg: "用户不存在", Code: 1000}
	PasswordError       = Error{Msg: "用户密码不对", Code: 1001}
	UserNameRepeatError = Error{Msg: "用户名已经存在", Code: 1002}
	UnLoginError        = Error{Msg: "用户未登录", Code: 1003}

	EquipmentNotExistError    = Error{Msg: "设备不存在", Code: 2000}
	EquipmentOfflineError     = Error{Msg: "设备离线", Code: 2001}
	EquipmentNotBindError     = Error{Msg: "设备未绑定", Code: 2002}
	EquipmentAlreadyBindError = Error{Msg: "设备已绑定", Code: 2003}
	EquipmentBindError        = Error{Msg: "设备绑定失败", Code: 2004}
	EquipmentUnbindError      = Error{Msg: "设备解绑失败", Code: 2005}
	EquipmentCodeError        = Error{Msg: "设备码不对", Code: 2006}
	EquipmentCodeRepeatError  = Error{Msg: "设备码已存在", Code: 2007}
)

type Error struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}
