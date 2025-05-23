package util_test

import (
	"fmt"

	"github.com/wxw9868/util"
)

func ExampleVerifyPassword() {
	err := util.VerifyPassword("99999.")
	fmt.Println(err)
}

func ExampleVerifyPayPassword() {
	fmt.Println(util.VerifyPayPassword("123456"))
}

func ExampleVerifyEmail() {
	fmt.Println(util.VerifyEmail("1234@163.com"))
}

func ExampleVerifyMobile() {
	fmt.Println(util.VerifyMobile("18201108888"))
}

func ExampleVerifyTelephone() {
	fmt.Println(util.VerifyTelephone("028-02866250077"))
}

func ExampleVerifyString() {
	fmt.Println(util.VerifyString("18201108888"))
}

func ExampleFormatTimeToUnix() {
	fmt.Println(util.FormatTimeToUnix("2006-01-02 15:04:05"))
}

func ExampleVerifyEnglish() {
	fmt.Println(util.VerifyEnglish("we"))
}

func ExampleVerifyFloat2f() {
	fmt.Println(util.VerifyFloat2f("1.11"))
}

func ExampleStringToTime() {
	fmt.Println(util.StringToTime("Asia/Shanghai", "2021-09-30 15:58:17"))
}

func ExampleGenerateCode() {
	for i := 0; i < 10000; i++ {
		code := util.GenerateCode(6)
		fmt.Println(code)
	}
}

func ExampleStringBuilder() {
	ss := []string{"a", "b", "c"}
	fmt.Println(util.StringBuilder(ss...))
}
