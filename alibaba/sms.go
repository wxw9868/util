package alibaba

import (
	"errors"
	"fmt"
	"time"

	"part-time/user-service/config"
	"part-time/user-service/pkg/redis"
	"part-time/user-service/pkg/util/tool"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

//发送验阿里云证码
func SendSMS(mobile string) error {
	var configSMS = config.GetConfig().AliSms
	client, err := dysmsapi.NewClientWithAccessKey(configSMS.RegionId, configSMS.AccessKeyId, configSMS.AccessKeySecret)
	if err != nil {
		return err
	}

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = mobile
	request.SignName = configSMS.SignName
	request.TemplateCode = configSMS.TemplateCode

	resCode := tool.GenerateCode(6)
	err = redis.GetRedisClient().Set(mobile, resCode, 15*time.Minute).Err()
	if err != nil {
		return err
	}

	request.TemplateParam = "{\"code\":" + resCode + "}"
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
