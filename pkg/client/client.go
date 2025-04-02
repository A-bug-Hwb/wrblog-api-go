package client

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"net"
	"wrblog-api-go/app/common/token"
	"wrblog-api-go/pkg/utils"
	"wrblog-api-go/pkg/utils/curl"
)

// ip地址
type IpAddress struct {
	Ip         string `json:"ip"`
	Pro        string `json:"pro"`
	ProCode    string `json:"proCode"`
	City       string `json:"city"`
	CityCode   string `json:"cityCode"`
	Region     string `json:"region"`
	RegionCode string `json:"regionCode"`
	Addr       string `json:"addr"`
	Browser    string `json:"browser"`
	Os         string `json:"os"`
}

func GetClient(c *gin.Context) (clientInfo *token.ClientInfo) {
	ipAddress := GetAddress(c.ClientIP(), c.Request.UserAgent())
	clientInfo = &token.ClientInfo{
		IpAddr:        ipAddress.Ip,
		LoginLocation: ipAddress.Addr,
		Browser:       ipAddress.Browser,
		Os:            ipAddress.Os,
	}
	return
}

// 根据ip获取地址
func GetAddress(ip string, userAgent string) (ipAddress *IpAddress) {
	// 解析userAgent
	userAgentData := useragent.Parse(userAgent)
	ipAddress = &IpAddress{
		Browser: userAgentData.Name,
		Os:      userAgentData.OS,
	}
	var internalIp = "(((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))\\.){3}((\\d)|([1-9]\\d)|(1\\d{2})|(2[0-4]\\d)|(25[0-5]))$"
	if utils.CheckRegex(internalIp, ip) || ip == "127.0.0.1" || ip == "::1" {
		ipAddress.Ip = ip
		ipAddress.Addr = "内网地址"
		return
	}
	if netIp := net.ParseIP(ip); netIp == nil || netIp.IsLoopback() {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return
	}
	body, err := curl.DefaultClient().Send(&curl.RequestParam{
		Url: "http://whois.pconline.com.cn/ipJson.jsp",
		Query: map[string]interface{}{
			"ip":   ip,
			"json": true,
		},
	})
	if err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return
	}
	if err = json.Unmarshal([]byte(body), &ipAddress); err != nil {
		ipAddress.Ip = ip
		ipAddress.Addr = "未知地址"
		return
	}
	return
}
