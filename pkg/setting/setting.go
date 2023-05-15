package setting

import (
	"fmt"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/klog/v2"

	"github.com/xishengcai/cloud/pkg/common"
	"github.com/xishengcai/cloud/pkg/file"
)

var (
	configFilePaths = [3]string{"conf/config.yaml", "../conf/config.yaml", "../../conf/config.yaml"}
	EnvConfig       *envConfig
	AliCloud        aliCloud
	Web             web
)

type envConfig struct {
	Title      string                  `yaml:"title"`
	ReleaseEnv string                  `yaml:"releaseEnv"`
	Version    string                  `yaml:"version"`
	Server     map[string]ServerConfig `yaml:"server"`
	RunMode    string                  `yaml:"runMode"`
}

type ServerConfig struct {
	Web      web      `yaml:"web"`
	AliCloud aliCloud `yaml:"aliCloud"`
}

type web struct {
	Port         string        `yaml:"port"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
	RunMode      string        `yaml:"runMode"`
}

type aliCloud struct {
	AK  string `yaml:"ak"`
	SK  string `yaml:"sk"`
	OSS oss    `yaml:"oss"`
}

type oss struct {
	AK       string `yaml:"ak"`
	SK       string `yaml:"sk"`
	Bucket   string `yaml:"bucket"`
	Endpoint string `yaml:"endpoint"`
	Domain   string `yaml:"domain"`
}

func init() {
	loadConfig("")
}

func loadConfig(path string) {
	if path != "" {
		// 根据传参初始化环境配置
		if file.NotExists(path) {
			panic(fmt.Errorf("config file path: %s not exists", path))
		}

		loadConfigWithPath(path)
		return
	}
	for _, configPath := range configFilePaths {
		// 从不同的层级目录初始化环境配置，直到有一次初始化成功后退出
		currentDir, _ := os.Getwd()
		klog.Infof("current directory: %s, load config: %s", currentDir, configPath)
		if file.NotExists(configPath) {
			continue
		}
		loadConfigWithPath(configPath)
		break
	}
	if EnvConfig == nil {
		panic("envConfig init fail")
	}
}

func loadConfigWithPath(configPath string) {
	config, err := os.ReadFile(configPath)
	if err != nil {
		klog.Error("read config file err: ", err)
		panic(err)
	}

	err = yaml.Unmarshal(config, &EnvConfig)
	if err != nil {
		panic(err)
	}
	klog.Infof("read Config: %s", configPath)

	releaseEnv := EnvConfig.ReleaseEnv
	serverConfig := EnvConfig.Server[releaseEnv]
	Web = serverConfig.Web
	AliCloud = serverConfig.AliCloud
	klog.Infof("config: %+v", common.PrettifyJson(EnvConfig.Server[releaseEnv], true))
}
