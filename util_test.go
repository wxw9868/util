/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 16:49:11
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-10 18:01:34
 * @FilePath: /util/util_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"fmt"
	"testing"
)

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

func TestVerifyFloat(t *testing.T) {
	fmt.Println(VerifyFloat("1.11"))
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
