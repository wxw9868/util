package util

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/skip2/go-qrcode"
)

// VerifyPassword 密码必须由字⺟、数字和_~!.@#$%^&*?-符号组成，长度为6 ~ 20个字符
// pattern := `^[\d|a-zA-Z]+[\d|a-zA-Z]+[_~!.@#$%^&*?-]+$`
//
//	if len(str) < 6 || len(str) > 20 {
//		return errors.New("密码长度为6 ~ 20个字符")
//	}
func VerifyPassword(str string) error {
	if len(str) < 6 || len(str) > 20 {
		return errors.New("密码长度为6 ~ 20个字符")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_~!.@#$%^&*?-]{6,20}$`).MatchString(str) {
		return errors.New("密码由字⺟、数字和符号（_~!.@#$%^&*?-）组成，长度为6 ~ 20个字符")
	}
	return nil
}

// VerifyPayPassword 支付密码验证规则:6个数字
func VerifyPayPassword(str string) error {
	if !regexp.MustCompile(`^\d{6}$`).MatchString(str) {
		return errors.New("支付密码由长度为6数字组成")
	}
	return nil
}

// VerifyEmail 验证邮箱
func VerifyEmail(str string) error {
	if !regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(str) {
		return errors.New("邮箱格式不正确")
	}
	return nil
}

// VerifyMobile 验证手机号
func VerifyMobile(str string) error {
	if !regexp.MustCompile(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9]|17[0-9])\d{8}$`).MatchString(str) {
		return errors.New("手机号格式不正确")
	}
	return nil
}

// VerifyTelephone 验证座机电话号
func VerifyTelephone(str string) error {
	if !regexp.MustCompile(`^((0\\d{2,3})-)(\\d{7,8})(-(\\d{3,}))?$`).MatchString(str) {
		return errors.New("座机电话号格式不正确")
	}
	return nil
}

// VerifyFloat 验证浮点数最多有两位小数
func VerifyFloat2f(f interface{}) error {
	if !regexp.MustCompile(`^\d+.?\d{0,2}$`).MatchString(fmt.Sprint(f)) {
		return errors.New("请输入正确的浮点数")
	}
	return nil
}

// VerifyChinese 匹配中文的匹配规则
func VerifyChinese(name string) error {
	if !regexp.MustCompile("^[\u4e00-\u9fa5]{2,10}").MatchString(name) {
		return errors.New("请输入正确的中文")
	}
	return nil
}

// VerifyEnglish 匹配英文的匹配规则
func VerifyEnglish(name string) error {
	if !regexp.MustCompile("^[a-zA-Z]{3,30}").MatchString(name) {
		return errors.New("请输入正确的英文")
	}
	return nil
}

// VerifyString 验证字符串
func VerifyString(str string) error {
	if !regexp.MustCompile("^[A-Za-z0-9]+$").MatchString(str) {
		return errors.New("字符串只能包含数字和英文")
	}
	return nil
}

// StringBuilder 拼接字符串
func StringBuilder(ss ...string) string {
	var builder strings.Builder
	for i := 0; i < len(ss); i++ {
		builder.WriteString(ss[i])
	}
	return builder.String()
}

// GenerateCode 生成6位数字码
func GenerateCode(width int) string {
	numeric := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var b strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&b, "%d", numeric[rand.Intn(r)])
	}
	return b.String()
}

// StringToTime 获取对应时区的实时时间
func StringToTime(timezone, value string) time.Time {
	if timezone == "" {
		timezone = "Asia/Shanghai"
	}
	loc, _ := time.LoadLocation(timezone)
	const longForm = "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(longForm, value, loc)
	return t
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// StartTime 指定时间
func StartTime(tt string) (d time.Time) {
	year, month, day := time.Now().Date()

	// 今日日期
	today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	switch tt {
	case "today":
		d = today
	case "yesterday":
		// 昨日日期
		d = today.AddDate(0, 0, -1)
	case "weekStart":
		// 本周起始日期（周一）
		d = today.AddDate(0, 0, -int(today.Weekday())+1)
	case "monthStart":
		// 本月起始日期
		d = time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	default:
		d = time.Now().Local()
	}
	return
}

// FormatTimeToUnix 格式化时间
func FormatTimeToUnix(tt string) int64 {
	const longForm = "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("UTC")
	t, _ := time.ParseInLocation(longForm, tt, loc)
	return t.Unix()
}

// 自增ID
var autoincrementID = uint64(rand.New(rand.NewSource(time.Now().Unix())).Int63n(10000))
var mutex sync.Mutex

// GenerateOrderSN 订单ID
func GenerateOrderSN() string {
	mutex.Lock()
	id := atomic.AddUint64(&autoincrementID, 1)
	if id > 9999 {
		autoincrementID = 0
	}
	mutex.Unlock()
	return fmt.Sprintf("%s%04d", time.Now().Format("20060102"), id)
}

// GenerateQRCode 生成二维码
func GenerateQRCode(url string, recoveryLevel string, size int, filename string) error {
	var level qrcode.RecoveryLevel
	switch recoveryLevel {
	case "low":
		level = qrcode.Low
	case "medium":
		level = qrcode.Medium
	case "high":
		level = qrcode.High
	case "highest":
		level = qrcode.Highest
	default:
		level = qrcode.Medium
	}
	if size == 0 {
		size = 256
	}
	if err := qrcode.WriteFile(url, level, size, filename); err != nil {
		return err
	}
	return nil
}
