package sms

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliyunSMS struct {
	RegionID        string
	AccessKeyID     string
	AccessKeySecret string
	SignName        string
	TemplateCode    string
}

func NewAliyunSMS(regionID, accessKeyID, accessKeySecret, signName, templateCode string) *AliyunSMS {
	return &AliyunSMS{
		RegionID:        regionID,
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
		SignName:        signName,
		TemplateCode:    templateCode,
	}
}

// SendSMS 发送验阿里云证码
func (sms *AliyunSMS) SendSMS(mobile string, code int) error {
	client, err := dysmsapi.NewClientWithAccessKey(sms.RegionID, sms.AccessKeyID, sms.AccessKeySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = sms.SignName
	request.TemplateCode = sms.TemplateCode
	request.TemplateParam = "{\"code\":" + strconv.Itoa(code) + "}"
	response, err := client.SendSms(request)
	if err != nil {
		return fmt.Errorf("短信发送失败：%s", err)
	}
	if response.Code != "OK" || response.Message != "OK" {
		return errors.New("短信发送失败！")
	}
	return nil
}
