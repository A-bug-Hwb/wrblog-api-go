package intercept

import (
	"regexp"
	"wrblog-api-go/config"
)

var noPath = []string{"/", "/favicon.ico", "/api/v1/demo/*", "/api/v1/auth/*", "/api/v1/file/*", "/swagger/*", "*/open/*", config.Conf.ConfigInfo.FilePrefix + "*"}

// NotIp 不拦截的ip(ip白名单)
func NotIp(ip string) bool {
	arr := config.Conf.ConfigInfo.IpWhites
	if len(arr) == 0 {
		return true
	}
	//默认本机不拦截
	arr = append(arr, "127.0.0.1", "::1")
	for _, element := range arr {
		if ip == element || element == "0.0.0.0" {
			return true
		}
	}
	return false
}

// NotIntercept 判断路由是否拦截
func NotIntercept(path string) bool {
	//不拦截的地址
	for _, url := range noPath {
		regexPattern := convertToRegex(url)
		matched, err := regexp.MatchString(regexPattern, path)
		if err != nil {
			continue
		}
		if matched {
			return true
		}
	}
	return false
}

// 将URL通配符转换为正则表达式
func convertToRegex(pattern string) string {
	// 转换*为正则表达式的.*
	re := regexp.MustCompile(`\*+`)
	pattern = re.ReplaceAllString(pattern, ".*")
	// 在模式的开始和结束处加上^和$
	return "^" + pattern + "$"
}
