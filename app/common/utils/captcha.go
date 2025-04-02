package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// Math 配置参数
var (
	Height          = 70
	Width           = 240
	NoiseCount      = 0
	ShowLineOptions = base64Captcha.OptionShowSineLine
	BgColor         = &color.RGBA{
		R: 40,
		G: 30,
		B: 89,
		A: 29,
	}
	FontsStorage base64Captcha.FontsStorage
	Fonts        []string
)

func CreateCode() (string, string, string, error) {
	id, b64s, answer, err := mathCaptcha().Generate()
	return id, b64s, answer, err
}

func mathCaptcha() *base64Captcha.Captcha {
	driver := base64Captcha.NewDriverMath(Height, Width, NoiseCount, ShowLineOptions, BgColor, FontsStorage, Fonts)
	captcha := base64Captcha.NewCaptcha(driver, store)
	return captcha
}
