package util

import (
	"os"
	"testing"
)

func TestIsVideoFileExt(t *testing.T) {
	ok := IsVideoFileExt("mp4")
	t.Log(ok)
}

func TestSecondsToFormatTime(t *testing.T) {
	t.Log(SecondsToFormatTime(1000))
}

func TestGetMP4Duration(t *testing.T) {
	videoDir := "./assets/video"
	filePath := videoDir + "/" + "test.mp4"
	fil, _ := os.Open(filePath)
	duration, _ := GetMP4Duration(fil)
	t.Log(duration)
}
