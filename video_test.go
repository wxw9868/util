/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2025-01-13 14:27:32
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:26:27
 * @FilePath: /util/mp4_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
	t.Log(SecondsToFormatTime(10000))
}

func TestGetMP4Duration(t *testing.T) {
	videoDir := "./assets/video"
	filePath := videoDir + "/" + "test.mp4"
	fil, _ := os.Open(filePath)
	duration, _ := GetMP4Duration(fil)
	t.Log(duration)
}
