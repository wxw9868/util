package util_test

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wxw9868/util"
)

type RegisterRequest struct {
	Mobile   string `json:"mobile" validate:"required,number,len=11" label:"手机号"`
	UserType int    `json:"user_type" validate:"required,number,len=1" label:"用户类别"`
	Captcha
}

type Captcha struct {
	VerifyCode string `json:"verify_code" validate:"required,number,len=6" label:"验证码"`
}

func ExampleNewValidate() {
	engine := gin.Default()
	engine.POST("/register", register)
	_ = engine.Run()
}

func register(c *gin.Context) {
	r := new(RegisterRequest)
	if err := util.NewValidate("label").StructError(r); err != nil {
		c.JSON(http.StatusNotFound, util.Msg(false, 0, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, util.Msg(true, 1, "注册成功", nil))
}

func ExampleField() {
	myEmail := "wxw9868"
	if err := util.NewValidate("label").FieldError(myEmail, "required,email"); err != nil {
		fmt.Println(err)
	}
}

func ExampleGetValidateTrans() {
	v, trans, err := util.NewValidate("").GetValidateTrans()
	fmt.Println(v)
	fmt.Println(trans)
	fmt.Println(err)
}

func ExampleVideoFileMode() {
	util.VideoFileMode("mp4")
	// Output:
	// true
}

func ExampleNewValidate_Gin() {
	v := util.NewValidate("json")
	v.InitValidateGin()
	s := v.GinError(errors.New("error"))
	fmt.Println(s)
	myEmail := "wxw9868"
	f := v.GinVar(myEmail, "required,email")
	fmt.Println(f)
}
