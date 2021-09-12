package common

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_cn "github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
)

// 数据验证器
func getValidate() (validate *validator.Validate, trans ut.Translator) {
	z := zh.New()
	uni := ut.New(z, z)
	trans, ok := uni.GetTranslator("zh")
	if !ok {
		log.Error("中文翻译器初始化失败！")
		return nil, nil
	}
	validate = validator.New()
	//注册一个函数，获取struct tag里自定义的label作为字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("label")
	})

	// 注册翻译器
	err := zh_cn.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Errorf("注册翻译器失败：%s", err)
	}
	return
}

var (
	validate *validator.Validate
	trans    ut.Translator
)

func Validate() *validator.Validate {
	if validate != nil {
		return validate
	}
	panic("数据验证器初始化失败！")
	return nil
}

func Trans() ut.Translator {
	if trans != nil {
		return trans
	}
	panic("字段翻译器初始化失败！")
	return nil
}

func init() {
	validate, trans = getValidate()
}
