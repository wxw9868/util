package util

import (
	"fmt"
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
		c.JSON(http.StatusOK, gin.H{
			"data":   nil,
			"code":   0,
			"status": false,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":   nil,
		"code":   1,
		"status": true,
		"msg":    "注册成功",
	})
	return
}

func TestGetValidate(t *testing.T) {
	engine := gin.Default()
	engine.POST("/register", register)
	_ = engine.Run()
}

func TestField(t *testing.T) {
	myEmail := "wxw9868"
	if err := NewValidate("label").FieldError(myEmail, "required,email"); err != nil {
		fmt.Println(err)
	}
}

func TestGetValidateTrans(t *testing.T) {
	v, trans, err := NewValidate("").GetValidateTrans()
	fmt.Println(v)
	fmt.Println(trans)
	fmt.Println(err)
}
