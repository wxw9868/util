package util

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
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

// VerifyPassword 登录密码必须由⼤⼩写字⺟、数字和符号，长度为6~20位
func VerifyPassword(str string) error {
	if len(str) < 6 {
		return errors.New("password len is < 9")
	}
	if len(str) > 20 {
		return errors.New("password len is not > 20")
	}
	num := `[0-9]{1}`
	aZ := `[a-z]{1}`
	AZ := `[A-Z]{1}`
	symbol := `[.!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, str); !b || err != nil {
		return errors.New("password need num :" + err.Error())
	}
	if b, err := regexp.MatchString(aZ, str); !b || err != nil {
		return errors.New("password need a_z :" + err.Error())
	}
	if b, err := regexp.MatchString(AZ, str); !b || err != nil {
		return errors.New("password need A_Z :" + err.Error())
	}
	if b, err := regexp.MatchString(symbol, str); !b || err != nil {
		return errors.New("password need symbol :" + err.Error())
	}
	return nil
}

// VerifyPayPassword 支付密码验证规则:6个数字
func VerifyPayPassword(str string) (bool, error) {
	matched, err := regexp.MatchString("^\\d{6}$", str)
	if err != nil {
		return false, err
	}
	return matched, nil
}

// VerifyEmail 验证邮箱
func VerifyEmail(str string) bool {
	matched, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", str)
	return matched
}

// VerifyMobile 验证手机号
func VerifyMobile(str string) bool {
	matched, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9]|17[0-9])\\d{8}$", str)
	return matched
}

// VerifyEnglish 验证字符串是否全为英文
func VerifyEnglish(str string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z]+$", str)
	return matched
}

// VerifyFloat2f 验证浮点数最多有两位小数
func VerifyFloat2f(str string) bool {
	matched, _ := regexp.MatchString("^(([1-9]{1}\\d*)|([0]{1}))(\\.(\\d){0,2})?$", str)
	return matched
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
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var b strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&b, "%d", numeric[rand.Intn(r)])
	}
	var s strings.Builder
	array := strings.Split(b.String(), "")
	firstCharacter, _ := strconv.Atoi(array[0])
	if firstCharacter == 0 {
		s.WriteString("8")
		s.WriteString(strings.Join(array[1:], ""))
		return s.String()
	}
	return b.String()
}
