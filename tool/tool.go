package tool

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"

	"github.com/rs/xid"
)

//数据加密
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

//登录密码必须由⼤⼩写字⺟、数字和符号，长度为6~20位
func VerifyLoginPassword(str string) error {
	if len(str) < 6 {
		return errors.New("password len is < 9")
	}
	if len(str) > 20 {
		return errors.New("password len is not > 20")
	}
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[.!@#~$%^&*()+|_]{1}`
	if b, err := regexp.MatchString(num, str); !b || err != nil {
		return errors.New("password need num :" + err.Error())
	}
	if b, err := regexp.MatchString(a_z, str); !b || err != nil {
		return errors.New("password need a_z :" + err.Error())
	}
	if b, err := regexp.MatchString(A_Z, str); !b || err != nil {
		return errors.New("password need A_Z :" + err.Error())
	}
	if b, err := regexp.MatchString(symbol, str); !b || err != nil {
		return errors.New("password need symbol :" + err.Error())
	}
	return nil
}

//支付密码验证规则:6个数字
func VerifyPayPassword(str string) bool {
	matched, _ := regexp.MatchString("^\\d{6}$", str)
	return matched
}

//验证邮箱
func VerifyEmail(str string) bool {
	matched, _ := regexp.MatchString("^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$", str)
	return matched
}

//验证手机号
func VerifyMobile(str string) bool {
	matched, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9]|17[0-9])\\d{8}$", str)
	return matched
}

//验证字符串是否全为英文
func VerifyEnglish(str string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z]+$", str)
	return matched
}

//拼接字符串
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

//生成应用AppId
func GetAppId() string {
	guid := xid.New()
	return guid.String()
}

//验证浮点数最多有两位小数
func VerifyFloat2f(str string) bool {
	matched, _ := regexp.MatchString("^(([1-9]{1}\\d*)|([0]{1}))(\\.(\\d){0,2})?$", str)
	return matched
}

//将用符号链接的英文字符的首字符转换为大写
func INITCAP(s, sep string) string {
	list := strings.Split(s, sep)
	var str string
	for j := 0; j < len(list); j++ {
		str += strings.Title(list[j])
	}
	return str
}

//获取对应时区的实时时间
func GetNowTime(timezone string) time.Time {
	if timezone == "" {
		timezone = "Asia/Shanghai"
	}
	loc, _ := time.LoadLocation(timezone)
	const longForm = "2006-01-02 15:04:05"
	t, _ := time.ParseInLocation(longForm, time.Now().UTC().String(), loc)
	return t
}

//生成6位数字码
func GenerateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var b strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&b, "%d", numeric[rand.Intn(r)])
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

//PKCS7 填充模式
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}
