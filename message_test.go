/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 20:05:23
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:39:55
 * @FilePath: /util/message_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"testing"
)

func TestSuccess(t *testing.T) {
	res := Success("成功", nil)
	t.Log(res)
}

func TestFail(t *testing.T) {
	res := Fail("失败")
	t.Log(res)
}

func TestMsg(t *testing.T) {
	res := Msg(1, "成功", nil)
	t.Log(res)
}

func TestCodeMsg(t *testing.T) {
	res := CodeMsg(-1, nil)
	t.Log(res)
}
