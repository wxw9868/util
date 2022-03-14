package util

//gin > 1.4.0
//将验证器错误翻译成中文

import (
	"reflect"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
)

func (v *Validate) initValidateGin() {
	v.trans, _ = ut.New(zh.New()).GetTranslator("zh")
	v.validate = binding.Validator.Engine().(*validator.Validate)
	//注册一个函数，获取struct tag里自定义的label作为字段名
	v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(v.tag)
	})
	_ = translation.RegisterDefaultTranslations(v.validate, v.trans)
}

func (v *Validate) GinError(err error) (ret string) {
	if validationErrors, ok := err.(validator.ValidationErrors); !ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			ret += e.Translate(v.trans) + ";"
		}
	}
	return ret
}

func (v *Validate) GinVar(f interface{}, t string) error {
	err := v.validate.Var(f, t)
	if err != nil {
		return err
	}
	return nil
}
