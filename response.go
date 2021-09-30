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

// Failure 默认失败
func Failure(msg string) Message {
	return Msg(false, CodeFail, msg, nil)
}

// Msg 序列化消息
func Msg(status bool, code int, msg string, data interface{}) Message {
	if msg == "" {
		if val, ok := MsgCodeMap[code]; ok {
			msg = val
		} else {
			msg = MsgCodeMap[-1]
		}
	}
	m := Message{
		Code:   code,
		Status: status,
		Msg:    msg,
		Data:   data,
	}
	return m
}

func GetCodeMsg(code int) string {
	if val, ok := MsgCodeMap[code]; ok {
		return val
	}
	return MsgCodeMap[-1]
}
