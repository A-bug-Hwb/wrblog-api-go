package utils

import "regexp"

func VerifyMobile(mobile string) bool {
	// 正则表达式匹配手机号格式
	regex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return regex.MatchString(mobile)
}

func VerifyMailBox(mailbox string) bool {
	// 正则表达式匹配邮箱格式
	regex := regexp.MustCompile(`^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`)
	return regex.MatchString(mailbox)
}

func VerifySpaceUrl(spaceUrl string) bool {
	// 正则表达式匹配空间地址格式
	regex := regexp.MustCompile(`^[a-zA-Z][0-9a-zA-Z_-]{2,35}$`)
	//spaceUrlRegex := regexp.MustCompile(`^(?!-)((?!--)[0-9a-zA-Z-]){1,39}(?<!-)$`)
	return regex.MatchString(spaceUrl)
}
