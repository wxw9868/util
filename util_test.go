package util

import (
	"testing"
	"time"
)

func TestVerifyPassword(t *testing.T) {
	err := VerifyPassword("123456")
	t.Log(err)
}

func TestVerifyPayPassword(t *testing.T) {
	err := VerifyPayPassword("123456")
	t.Log(err)
}

func TestVerifyEmail(t *testing.T) {
	err := VerifyEmail("1234@163.com")
	t.Log(err)
}

func TestVerifyMobile(t *testing.T) {
	err := VerifyMobile("18201108888")
	t.Log(err)
}

func TestVerifyTelephone(t *testing.T) {
	err := VerifyTelephone("028-02866250077")
	t.Log(err)
}

func TestVerifyFloat2f(t *testing.T) {
	err := VerifyFloat2f("1.111")
	t.Log(err)
}

func TestVerifyChinese(t *testing.T) {
	err := VerifyChinese("中国")
	t.Log(err)
}

func TestVerifyEnglish(t *testing.T) {
	err := VerifyEnglish("china")
	t.Log(err)
}

func TestVerifyString(t *testing.T) {
	err := VerifyString("中国....18201108888sdsaasa")
	t.Log(err)
}

func TestStringBuilder(t *testing.T) {
	t.Log(StringBuilder("a", "b", "c"))
}

func TestGenerateCode(t *testing.T) {
	code := GenerateCode(6)
	t.Log(code)
}

func TestStringToTime(t *testing.T) {
	t.Log(StringToTime("Asia/Shanghai", "2021-09-30 15:58:17"))
}

func TestFormatTime(t *testing.T) {
	t.Log(FormatTime(time.Now()))
}

func TestStartTime(t *testing.T) {
	t.Log(StartTime("yesterday"))
}

func TestFormatToUnix(t *testing.T) {
	t.Log(FormatTimeToUnix("2006-01-02 15:04:05"))
}

func TestGenerateOrderSN(t *testing.T) {
	t.Log(GenerateOrderSN())
}

func TestGenerateQRCode(t *testing.T) {
	err := GenerateQRCode("https://www.baidu.com/", "highest", 0, "qrcode.png")
	t.Log(err)
}
