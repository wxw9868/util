package util

import (
	"fmt"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	key := []byte("0123456789abcdef")
	res, err := AesEncrypt([]byte("123456"), key)
	if err != nil {
		fmt.Println(err)
	}
	res, err = AesDeCrypt(res, key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
}

func TestMd5(t *testing.T) {
	fmt.Println(Md5("123456"))
}
