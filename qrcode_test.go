package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	err := GenerateQRCode("https://www.baidu.com/", "highest", 0, "qrcode.png")
	fmt.Println(err)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
}
