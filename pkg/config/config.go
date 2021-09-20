package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"path"
	"runtime"
	"strings"
)

type Config struct {
	filePath string
	configs  map[string]interface{}
}

func NewConfig(filePaths ...string) *Config {
	c := &Config{
		configs: make(map[string]interface{}),
	}
	if len(filePaths) != 0 {
		c.filePath = filePaths[0]
	}

	return c
}

//注册配置
func (c *Config) Register(section string, config interface{}) *Config {
	c.configs[section] = config

	return c
}

func (c *Config) Parse() (err error) {

	//if err = c.loadFromFile(); err != nil {
	//	return
	//}

	if err = c.loadFromEnv(); err != nil {
		return
	}

	return
}

//配置文件路径
func (c *Config) getFilePath() string {
	if c.filePath != "" {
		return c.filePath
	}

	_, file, _, _ := runtime.Caller(1)
	file = file

	return ""
}

//从文件载入配置
func (c *Config) loadFromFile() (err error) {
	filePath := c.getFilePath()

	dir, file := path.Split(filePath)
	if arr := strings.Split(file, "."); len(arr) == 2 {
		viper.SetConfigName(arr[0])
		viper.SetConfigType(arr[1])
	}

	viper.AddConfigPath(dir)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	for k, v := range c.configs {
		if err = viper.UnmarshalKey(k, v); err != nil {
			return
		}
	}

	//watch 热更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		for k, v := range c.configs {
			if err = viper.UnmarshalKey(k, v); err != nil {
				return
			}
		}
	})
	viper.WatchConfig()

	return
}

//从环境变量载入配置
func (c *Config) loadFromEnv() (err error) {
	el := EnvironmentLoader{}
	for k, v := range c.configs {
		el.Prefix = strings.ToUpper(k)
		if err = el.Load(v); err != nil {
			return
		}
	}
	return
}

//载入默认配置
func (c *Config) loadFromDefault() (err error) {
	d := NewDefault()
	for _, v := range c.configs {
		if err = d.Load(v); err != nil {
			return
		}
	}
	return
}
