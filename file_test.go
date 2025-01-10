/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 16:49:10
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-10 18:08:23
 * @FilePath: /util/file_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestUploadedFile(t *testing.T) {
	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		fh, err := NewFile(request).FormFile("file")
		if err != nil {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("上传失败！"))
			return
		}
		dst := "./test.jpg"
		if err = SaveUploadedFile(fh, dst); err != nil {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("上传失败！"))
			return
		}
		writer.WriteHeader(http.StatusOK)
		data := map[string]interface{}{
			"data":   nil,
			"code":   1,
			"status": true,
			"msg":    "上传成功",
		}
		b, _ := json.Marshal(data)
		_, _ = writer.Write(b)
	})
	_ = http.ListenAndServe(":8080", nil)
}

func TestPathExists(t *testing.T) {
	path := "./test.jpg"
	ok := PathExists(path)
	if ok {
		_ = os.Remove(path) // 删除文件
	}
}
