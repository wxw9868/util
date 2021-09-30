package util

import (
	"github.com/skip2/go-qrcode"
)

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
