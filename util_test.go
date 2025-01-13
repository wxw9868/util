/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 16:49:11
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:01:49
 * @FilePath: /util/util_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"testing"
	"time"
)

func TestVerifyPassword(t *testing.T) {
	err := VerifyPassword("12345")
	t.Log(err)
}

func TestVerifyPayPassword(t *testing.T) {
	t.Log(VerifyPayPassword("123456"))
}

func TestVerifyEmail(t *testing.T) {
	t.Log(VerifyEmail("1234@163.com"))
}

func TestVerifyMobile(t *testing.T) {
	t.Log(VerifyMobile("18201108888"))
}

func TestVerifyTelephone(t *testing.T) {
	t.Log(VerifyTelephone("028-02866250077"))
}

func TestVerifyString(t *testing.T) {
	t.Log(VerifyString("18201108888"))
}

func TestVerifyNickname(t *testing.T) {
	t.Log(VerifyNickname("18201108888"))
}

func TestVerifyEnglish(t *testing.T) {
	t.Log(VerifyEnglish("we"))
}

func TestVerifyFloat2f(t *testing.T) {
	t.Log(VerifyFloat2f("1.11"))
}

func TestGetNowTime(t *testing.T) {
	t.Log(LocNowTime("Asia/Shanghai", "2021-09-30 15:58:17"))
}

func TestGenerateCode(t *testing.T) {
	for i := 0; i < 10000; i++ {
		code := GenerateCode(6)
		t.Log(code)
	}
}

func TestGetOrderID(t *testing.T) {
	t.Log(GenerateOrderSN())
}

func TestStringBuilder(t *testing.T) {
	t.Log(StringBuilder("a", "b"))
}

func TestINITCAP(t *testing.T) {
	t.Log(INITCAP("a", "b"))
}

func TestFormatToUnix(t *testing.T) {
	t.Log(FormatTimeToUnix("2006-01-02 15:04:05"))
}

func TestFormatTime(t *testing.T) {
	t.Log(time.Now())
}

func TestStartTime(t *testing.T) {
	t.Log(StartTime("yesterday"))
}
