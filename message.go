/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 16:49:10
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:30:44
 * @FilePath: /util/message.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
