package sms

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	err := NewAliyunSMS("", "", "", "", "").SendSMS("18200001111", 123456)
	if err != nil {
		fmt.Println(err)
	}
}
