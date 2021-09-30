package alibaba

import (
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	err := NewSMS("", "", "", "", "").SendSMS("18200001111", "123456")
	if err != nil {
		fmt.Println(err)
	}
}
