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
