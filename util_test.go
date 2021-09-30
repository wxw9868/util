package util

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestDataEncryption(t *testing.T) {
	str, err := DataEncryption("zh1234567")
	fmt.Println(str, err)
}

func TestVerifyPassword(t *testing.T) {
	err := VerifyPassword("zh1234567")
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
	count := 0
	for i := 0; i < 1000000; i++ {
		code := GenerateCode(6)
		array := strings.Split(code, "")
		firstCharacter, _ := strconv.Atoi(array[0])
		if firstCharacter == 0 {
			count++
		}
	}
	fmt.Println(count)
}
