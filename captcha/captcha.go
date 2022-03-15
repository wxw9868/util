package captcha

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

func GetCaptcha(captchaType string) map[string]interface{} {
	param := configJsonBody{CaptchaType: captchaType}
	switch captchaType {
	case "audio":
		param.DriverAudio = &base64Captcha.DriverAudio{
			Length:   6,
			Language: "zh",
		}
	case "string":
		param.DriverString = &base64Captcha.DriverString{
			Height:          60,
			Width:           240,
			NoiseCount:      0,
			ShowLineOptions: 0,
			Length:          6,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
			BgColor: &color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		}
	case "math":
		param.DriverMath = &base64Captcha.DriverMath{
			Height:          60,
			Width:           240,
			NoiseCount:      0,
			ShowLineOptions: 0,
			BgColor: &color.RGBA{
				R: 0,
				G: 0,
				B: 0,
				A: 0,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		}
	case "chinese":
		param.DriverChinese = &base64Captcha.DriverChinese{
			Height:          60,
			Width:           320,
			NoiseCount:      0,
			ShowLineOptions: 0,
			Length:          2,
			Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值",
			BgColor: &color.RGBA{
				R: 125,
				G: 125,
				B: 0,
				A: 118,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		}
	default:
		param.DriverDigit = &base64Captcha.DriverDigit{
			Height:   80,
			Width:    240,
			Length:   5,
			MaxSkew:  0.7,
			DotCount: 80,
		}
	}
	return GenerateCaptcha(param)
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha base64Captcha create http handler
func GenerateCaptcha(param configJsonBody) (body map[string]interface{}) {
	var driver base64Captcha.Driver
	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body = map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}
	return
}

func Verify(Id, VerifyValue string) bool {
	param := configJsonBody{
		Id:          Id,
		VerifyValue: VerifyValue,
	}
	//verify the captcha
	if store.Verify(param.Id, param.VerifyValue, true) {
		return true
	}
	return false
}
