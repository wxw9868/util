/*
 * @Author: wxw9868 wxw9868@163.com
 * @Date: 2024-01-16 20:03:58
 * @LastEditors: wxw9868 wxw9868@163.com
 * @LastEditTime: 2025-01-13 15:31:39
 * @FilePath: /util/example_validate_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util_test

import (
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
		c.JSON(http.StatusNotFound, util.Msg(0, err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, util.Msg(1, "注册成功", nil))
}
