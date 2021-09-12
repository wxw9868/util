package common

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// Context is the most important part of gin. It allows us to pass variables between middleware,
// manage the flow, validate the JSON of a request and render a JSON response for example.
type Context struct {
	Request            *http.Request
	MaxMultipartMemory int64
}

// FormFile returns the first file for the provided form key.
func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(c.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}

// SaveUploadedFile uploads the form file to specific dst.
func SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// 删除文件
func RemoveFile(filename string) error {
	return os.Remove(filename)
}

// 文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
