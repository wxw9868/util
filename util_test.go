package util

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password, err := NewPassword("wxw9868").Encrypt("123456")
	fmt.Printf("password: %s, err: %s\n", password, err)
}

func TestVerifyPassword(t *testing.T) {
	err := VerifyPassword("123456*")
	fmt.Println(err)
}

func TestVerifyPayPassword(t *testing.T) {
	fmt.Println(VerifyPayPassword("123456"))
}

func TestVerifyEmail(t *testing.T) {
	fmt.Println(VerifyEmail("1234@163.com"))
}

func TestVerifyMobile(t *testing.T) {
	fmt.Println(VerifyMobile("18201108888"))
}

func TestVerifyTelephone(t *testing.T) {
	fmt.Println(VerifyTelephone("028-02866250077"))
}

func TestVerifyString(t *testing.T) {
	fmt.Println(VerifyString("18201108888"))
}

func TestVerifyName(t *testing.T) {
	fmt.Println(VerifyName("18201108888"))
}

func TestFormatToUnix(t *testing.T) {
	fmt.Println(FormatToUnix("2006-01-02 15:04:05"))
}

func TestVerifyEnglish(t *testing.T) {
	fmt.Println(VerifyEnglish("we"))
}

func TestVerifyFloat2f(t *testing.T) {
	fmt.Println(VerifyFloat2f("1.11"))
}

func TestGetNowTime(t *testing.T) {
	fmt.Println(GetNowTime("Asia/Shanghai", "2021-09-30 15:58:17"))
}

func TestGenerateCode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		code := GenerateCode(6)
		fmt.Println(code)
	}
}

func TestGetOrderID(t *testing.T) {
	fmt.Println(GenerateOrderSN())
}

func TestStringBuilder(t *testing.T) {
	fmt.Println(StringBuilder("a", "b"))
}

func TestINITCAP(t *testing.T) {
	fmt.Println(INITCAP("a", "b"))
}
