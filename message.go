package util

const (
	CodeSuccess = 1
	CodeFail    = 0
	CodeError   = -1
)

var MsgCodeMap = map[int]string{
	CodeSuccess: "success",
	CodeFail:    "fail",
	CodeError:   "unknown error",
}

type Message struct {
	Code   int         `json:"code"`
	Status bool        `json:"status"`
	Msg    string      `json:"message"`
	Data   interface{} `json:"data"`
}

// Success 默认成功
func Success(msg string, data interface{}) Message {
	return Msg(true, CodeSuccess, msg, data)
}

// Fail 默认失败
func Fail(msg string) Message {
	return Msg(false, CodeFail, msg, nil)
}

// Msg 序列化消息
func Msg(status bool, code int, msg string, data interface{}) Message {
	return Message{code, status, msg, data}
}

// CodeMsg 匹配状态码和信息
func CodeMsg(status bool, code int, data interface{}) Message {
	return Message{code, status, MsgCodeMap[code], data}
}
