package util

import (
	"bytes"
	"errors"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhcn "github.com/go-playground/validator/v10/translations/zh"
)

type Validate struct {
	tag      string
	validate *validator.Validate
	trans    ut.Translator
}

// NewValidate 初始化 Validate
func NewValidate(tag string) *Validate {
	return &Validate{
		tag: tag,
	}
}

// Validate 数据验证器
func (v *Validate) initValidate(tag string) error {
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
	return zhcn.RegisterDefaultTranslations(v.validate, v.trans)
}

// GetValidateTrans 获取配置
func (v *Validate) GetValidateTrans() (*validator.Validate, ut.Translator, error) {
	if err := v.initValidate(v.tag); err != nil {
		return nil, nil, err
	}
	return v.validate, v.trans, nil
}

// StructError 结构体验证
func (v *Validate) StructError(s interface{}) error {
	if err := v.initValidate(v.tag); err != nil {
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

// FieldError 字段验证
func (v *Validate) FieldError(field interface{}, tag string) error {
	if err := v.initValidate(v.tag); err != nil {
		return err
	}
	if err := v.validate.Var(field, tag); err != nil {
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
