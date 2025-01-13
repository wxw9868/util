/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 19:52:13
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:41:33
 * @FilePath: /util/validate_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Mobile   string `json:"mobile" validate:"required,number,len=11" label:"手机号"`
	UserType int    `json:"user_type" validate:"required,number,len=1" label:"用户类别"`
	Captcha
}

type Captcha struct {
	VerifyCode string `json:"verify_code" validate:"required,number,len=6" label:"验证码"`
}

func register(c *gin.Context) {
	r := new(RegisterRequest)
	if err := NewValidate("label").StructError(r); err != nil {
		c.JSON(http.StatusNotFound, Msg(0, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, Msg(1, "注册成功", nil))
}

func TestNewValidate(t *testing.T) {
	engine := gin.Default()
	engine.POST("/register", register)
	_ = engine.Run()
}

func TestField(t *testing.T) {
	myEmail := "wxw9868"
	if err := NewValidate("label").FieldError(myEmail, "required,email"); err != nil {
		t.Log(err)
	}
}

func TestGetValidateTrans(t *testing.T) {
	v, trans, err := NewValidate("").GetValidateTrans()
	t.Log(v, trans, err)
}
