package util

import (
	"fmt"
	"testing"
)

func TestDataEncryption(t *testing.T) {
	str, err := DataEncryption("zh1234567")
	fmt.Println(str, err)
}

func TestVerifyPassword(t *testing.T) {
	err := VerifyPassword("99999.")
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
