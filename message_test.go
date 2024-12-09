package util

import (
	"fmt"
	"testing"
)

func TestSuccess(t *testing.T) {
	res := Success("成功", nil)
	fmt.Println(res)
}

func TestFail(t *testing.T) {
	res := Fail("失败")
	fmt.Println(res)
}

func TestMsg(t *testing.T) {
	Msg(true, 1, "成功", nil)
}

func TestCodeMsg(t *testing.T) {
	res := CodeMsg(false, -1, nil)
	fmt.Println(res)
}
