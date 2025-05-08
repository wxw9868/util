package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetIp 获取客户端ip
func GetIp() (string, error) {
	conn, err := net.Dial("udp", "google.com:80")
	if err != nil {
		return "", nil
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}

// GetExternalIP 通过http://myexternalip.com/raw获取公网IP
func GetExternalIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return GetExternalIP()
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	content, _ := io.ReadAll(resp.Body)
	return string(content)
}

// GetPublicIP 通过dns服务器8.8.8.8:80获取使用的ip
func GetPublicIP() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	return localAddr[0:idx]
}

// GetIntranetIp 获取本地ip
func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println("ip:", ipnet.IP.String())
			}
		}
	}
}

// IsPublicIP 判断是否是公网ip
func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// ip地址string转int
func inetAton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")
	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])
	var sum int64
	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)
	return sum
}

// ip地址int转string
func inetNtoa(ipnr int64) net.IP {
	var b [4]byte
	b[0] = byte(ipnr & 0xFF)
	b[1] = byte((ipnr >> 8) & 0xFF)
	b[2] = byte((ipnr >> 16) & 0xFF)
	b[3] = byte((ipnr >> 24) & 0xFF)
	return net.IPv4(b[3], b[2], b[1], b[0])
}

// IpBetween 判断ip地址区间
func IpBetween(from net.IP, to net.IP, test net.IP) bool {
	if from == nil || to == nil || test == nil {
		fmt.Println("An ip input is nil") // or return an error!? return false
	}
	//goland:noinspection GoDfaNilDereference
	from16 := from.To16()
	//goland:noinspection GoDfaNilDereference
	to16 := to.To16()
	//goland:noinspection GoDfaNilDereference
	test16 := test.To16()
	if from16 == nil || to16 == nil || test16 == nil {
		fmt.Println("An ip did not convert to a 16 byte") // or return an error!?
		return false
	}
	if bytes.Compare(test16, from16) >= 0 && bytes.Compare(test16, to16) <= 0 {
		return true
	}
	return false
}

// IPInfo 通过淘宝接口根据公网ip获取国家运营商等信息
type IPInfo struct {
	Code int `json:"code"`
	Data IP  `json:"data"`
}

type IP struct {
	Country   string `json:"country"`
	CountryId string `json:"country_id"`
	Area      string `json:"area"`
	AreaId    string `json:"area_id"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	City      string `json:"city"`
	CityId    string `json:"city_id"`
	Isp       string `json:"isp"`
}

func TaobaoAPI(ip string) *IPInfo {
	url := "http://ip.taobao.com/service/getIpInfo.php?ip="
	url += ip
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	var result IPInfo
	if err = json.Unmarshal(out, &result); err != nil {
		return nil
	}
	return &result
}
