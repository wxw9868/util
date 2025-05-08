package util

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestAll(t *testing.T) {
	getIp, err := GetIp()
	fmt.Println(getIp, err)
	externalIp := GetExternalIP()
	externalIp = strings.Replace(externalIp, "\n", "", -1)
	fmt.Println("公网ip是: ", externalIp)
	fmt.Println("------Dividing Line------")
	ip := net.ParseIP(externalIp)
	if ip == nil {
		t.Error("您输入的不是有效的IP地址，请重新输入！")
	} else {
		result := TaobaoAPI(externalIp)
		if result != nil {
			fmt.Println("国家：", result.Data.Country)
			fmt.Println("地区：", result.Data.Area)
			fmt.Println("城市：", result.Data.City)
			fmt.Println("运营商：", result.Data.Isp)
		}
	}
	fmt.Println("------Dividing Line------")
	GetIntranetIp()
	fmt.Println("------Dividing Line------")
	ipInt := inetAton(net.ParseIP(externalIp))
	fmt.Println("Convert IPv4 address to decimal number(base 10) :", ipInt)
	ipResult := inetNtoa(ipInt)
	fmt.Println("Convert decimal number(base 10) to IPv4 address:", ipResult)
	fmt.Println("------Dividing Line------")
	isBetween := IpBetween(net.ParseIP("0.0.0.0"), net.ParseIP("255.255.255.255"), net.ParseIP(externalIp))
	fmt.Println("check result: ", isBetween)
	fmt.Println("------Dividing Line------")
	isPublicIp := IsPublicIP(net.ParseIP(externalIp))
	fmt.Println("It is public ip: ", isPublicIp)
	isPublicIp = IsPublicIP(net.ParseIP("169.254.85.131"))
	fmt.Println("It is public ip: ", isPublicIp)
	fmt.Println("------Dividing Line------")
	fmt.Println("aaa:", GetPublicIP())
}
