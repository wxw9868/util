package tool

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

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

func TestDataEncryption(t *testing.T) {
	fmt.Println(DataEncryption("zh1234567"))
}
