package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const ConfPath = "./config/config.yaml"

// Config 服务映射
type config struct {
	Server struct {
		Port  int    `yaml:"port"`
		Host  string `yaml:"host"`
		Https bool   `yaml:"https"`
		Mode  string `yaml:"mode"`
	} `yaml:"server"`
	ConfigInfo struct {
		FilePrefix  string   `yaml:"file-prefix"`
		Profile     string   `yaml:"profile"`
		FileMaxSize string   `yaml:"file-max-size"`
		IpWhites    []string `yaml:"ip-white" yaml:"ipWhites"`
		Token       struct {
			Header     string `yaml:"header"`
			Secret     string `yaml:"secret"`
			ExpireTime int    `yaml:"expire-time"`
		} `yaml:"token"`
		Captcha string `yaml:"captcha"`
		Key     struct {
			PublicKey  string `yaml:"public-key"`
			PrivateKey string `yaml:"private-key"`
		} `yaml:"key"`
	} `yaml:"config"`
}

var Conf *config

func init() {
	analysisConfig()
}

func analysisConfig() {
	// 打开 YAML 文件
	file, err := os.ReadFile(ConfPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse service configuration:%s", err.Error()))
	}
	err = yaml.Unmarshal(file, &Conf)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse service configuration:%s", err.Error()))
	}
}
