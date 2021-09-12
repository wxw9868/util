package alibaba

import (
	"errors"
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

var regionID = ""
var accessKeyID = ""
var accessKeySecret = ""
var signName = ""
var templateCode = ""

// SendSMS 发送验阿里云证码
func SendSMS(mobile string, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = signName
	request.TemplateCode = templateCode
	request.TemplateParam = "{\"code\":" + code + "}"
	response, err := client.SendSms(request)
	if err != nil {
		return fmt.Errorf("短信发送失败：%s", err)
	}
	if response.Code != "OK" || response.Message != "OK" {
		fmt.Println("message: ", response.Message, "code: ", response.Code)
		return errors.New("短信发送失败！")
	}
	return nil
}
