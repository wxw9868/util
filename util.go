package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/crypto/scrypt"
)

// DataEncryption 数据加密
func DataEncryption(password string) (string, error) {
	// DO NOT use this salt value; generate your own random salt. 8 bytes is
	// a good length.
	salt := []byte{0xc7, 0x29, 0xd2, 0x98, 0xb7, 0x7a, 0xcd, 0x7b}

	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(dk), nil
}

// VerifyPassword 密码必须由字⺟、数字和_~!.@#$%^&*?-符号组成，长度为6 ~ 20个字符
// pattern := `^[\d|a-zA-Z]+[\d|a-zA-Z]+[_~!.@#$%^&*?-]+$`
//
//	if len(str) < 6 || len(str) > 20 {
//		return errors.New("密码长度为6 ~ 20个字符")
//	}
func VerifyPassword(str string) error {
	pattern := `^[a-zA-Z0-9_~!.@#$%^&*?-]{6,20}$`
	if b, _ := regexp.MatchString(pattern, str); !b {
		return errors.New("密码由字⺟、数字和符号（_~!.@#$%^&*?-）组成，长度为6 ~ 20个字符")
	}
	return nil
}

// VerifyPayPassword 支付密码验证规则:6个数字
func VerifyPayPassword(str string) (bool, error) {
	matched, err := regexp.MatchString(`^\d{6}$`, str)
	if err != nil {
		return false, err
	}
	return matched, nil
}

// VerifyEmail 验证邮箱
func VerifyEmail(str string) bool {
	matched, _ := regexp.MatchString(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`, str)
	return matched
}

// VerifyMobile 验证手机号
func VerifyMobile(str string) bool {
	matched, _ := regexp.MatchString(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9]|17[0-9])\d{8}$`, str)
	return matched
}

// VerifyTelephone 座机号格式校验
func VerifyTelephone(str string) bool {
	pattern := "^((0\\d{2,3})-)(\\d{7,8})(-(\\d{3,}))?$" //比如：028-02866250077或0312-4295xxx的格式
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

// VerifyString 验证字符串
func VerifyString(sn string) bool {
	isSN := regexp.MustCompile("^[A-Za-z0-9]+$") //匹配数字和英文的匹配规则
	return isSN.MatchString(sn)
}

// VerifyEnglish 验证字符串是否全为英文
func VerifyEnglish(str string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, str)
	return matched
}

// VerifyFloat2f 验证浮点数最多有两位小数
func VerifyFloat2f(str string) bool {
	matched, _ := regexp.MatchString(`^(([1-9]{1}\d*)|([0]{1}))(\.(\d){0,2})?$`, str)
	return matched
}

// VerifyFloat 验证浮点数最多有两位小数
func VerifyFloat(f interface{}) bool {
	isNum := regexp.MustCompile(`^\d+.?\d{0,2}$`)
	return isNum.MatchString(fmt.Sprint(f))
}

func VerifyName(name string) bool {
	isChinese := regexp.MustCompile("^[\u4e00-\u9fa5]{2,10}") //匹配中文的匹配规则
	isEnglish := regexp.MustCompile("^[a-zA-Z]{3,30}")        //匹配英文的匹配规则
	//使用MatchString来将要匹配的字符串传到匹配规则中
	if isChinese.MatchString(name) || isEnglish.MatchString(name) {
		return true
	}
	return false
}

func FormatToUnix(str string) int64 {
	const longForm = "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("UTC")
	t, _ := time.ParseInLocation(longForm, str, loc)
	return t.Unix()
}

// GetNowTime 获取对应时区的实时时间
func GetNowTime(timezone, value string) time.Time {
	if timezone == "" {
		timezone = "Asia/Shanghai"
	}
	loc, _ := time.LoadLocation(timezone)
	const longForm = "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(longForm, value, loc)
	return t
}

// GenerateCode 生成6位数字码
func GenerateCode(width int) string {
	numeric := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&b, "%d", numeric[rand.Intn(r)])
	}
	return b.String()
}

// StringBuilder 拼接字符串
func StringBuilder(s1, s2 string) string {
	// strings.Builder的0值可以直接使用
	var builder strings.Builder

	// 向builder中写入字符/字符串
	builder.Write([]byte(s1))
	builder.WriteByte(' ')
	builder.WriteString(s2)

	// String() 方法获得拼接的字符串
	return builder.String()
}

// INITCAP 将用符号链接的英文字符的首字符转换为大写
func INITCAP(s, sep string) string {
	list := strings.Split(s, sep)
	var str string
	for j := 0; j < len(list); j++ {
		str += strings.Title(list[j])
	}
	return str
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

// VideoFileMode 视频文件类型
func VideoFileMode(ext string) bool {
	exts := map[string]struct{}{
		"mp4": {}, "swf": {}, "flv": {},
		"rm": {}, "ram": {}, "mov": {},
		"mpg": {}, "mpeg": {}, "wmv": {}, "avi": {},
	}
	if _, ok := exts[ext]; ok {
		return ok
	}
	return false
}
