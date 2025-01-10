/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 16:49:10
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-10 18:02:45
 * @FilePath: /util/encrypt_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"testing"
)

func TestPassword(t *testing.T) {
	password, err := NewPassword("wxw9868").Encrypt("123456")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("password: %s\n", password)
}

func TestAesEncrypt(t *testing.T) {
	key := []byte("0123456789abcdef")
	res, err := AesEncrypt([]byte("123456"), key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))

	res, err = AesDeCrypt(res, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(res))
}

func TestMd5(t *testing.T) {
	t.Log(Md5("123456"))
}
