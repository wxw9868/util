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
