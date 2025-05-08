package util

var MsgCodeMap = map[int]string{
	1:  "SUCCESS",
	0:  "FAIL",
	-1: "UNKNOWN ERROR",
}

type Message struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success 默认成功
func Success(msg string, data interface{}) Message {
	return Msg(1, msg, data)
}

// Fail 默认失败
func Fail(msg string) Message {
	return Msg(0, msg, nil)
}

// Msg 序列化消息
func Msg(code int, msg string, data interface{}) Message {
	return Message{code, msg, data}
}

// CodeMsg 匹配状态码和信息
func CodeMsg(code int, data interface{}) Message {
	return Message{code, MsgCodeMap[code], data}
}
