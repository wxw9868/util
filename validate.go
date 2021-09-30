package util

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhcn "github.com/go-playground/validator/v10/translations/zh"
)

type Validate struct {
	tag      string
	s        interface{}
	validate *validator.Validate
	trans    ut.Translator
}

func NewValidate(tag string) *Validate {
	return &Validate{
		tag: tag,
	}
}

// GetValidate 数据验证器
func (v *Validate) GetValidate(tag string) (err error) {
	var ok bool
	v.trans, ok = ut.New(zh.New(), zh.New()).GetTranslator("zh")
	if !ok {
		return errors.New("中文翻译器初始化失败！")
	}
	v.validate = validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get(tag)
	})

	// 注册翻译器
	err = zhcn.RegisterDefaultTranslations(v.validate, v.trans)
	if err != nil {
		return errors.New(fmt.Sprintf("注册翻译器失败：%s", err))
	}
	return
}

func (v *Validate) Error(s interface{}) error {
	if err := v.GetValidate(v.tag); err != nil {
		return err
	}
	if err := v.validate.Struct(s); err != nil {
		var buffer bytes.Buffer
		if validationErrors, ok := err.(validator.ValidationErrors); !ok {
			return err
		} else {
			for _, e := range validationErrors {
				buffer.WriteString(e.Translate(v.trans) + ";")
			}
		}
		return errors.New(buffer.String())
	}
	return nil
}
