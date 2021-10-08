package config

import (
	"fmt"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	Option func(c *Config)
	Config struct {
		filePath     string
		envPrefix    string //环境变量前缀
		envCamelCase bool   //环境变量大写下划线处理
		configs      map[string]interface{}
	}
)

//WithFilePath 设置文件路径
func WithFilePath(f string) Option {
	return func(c *Config) {
		c.filePath = f
	}
}

//WithEnvPrefix 设置环境变量前缀
func WithEnvPrefix(f string) Option {
	return func(c *Config) {
		c.envPrefix = f
	}
}

//WithEnvCamelCase 环境变量大写下划线处理
func WithEnvCamelCase(b bool) Option {
	return func(c *Config) {
		c.envCamelCase = b
	}
}

func NewConfig(options ...Option) *Config {
	c := &Config{
		configs: make(map[string]interface{}),
	}

	for _, option := range options {
		option(c)
	}

	return c
}

//注册配置
func (c *Config) Register(section string, config interface{}) *Config {
	c.configs[section] = config

	return c
}

func (c *Config) Parse() (err error) {

	if err = c.LoadFromDefault(); err != nil {
		return
	}

	if err = c.LoadFromFile(); err != nil {
		return
	}

	if err = c.LoadFromEnv(); err != nil {
		return
	}

	return
}

//从文件载入配置
func (c *Config) LoadFromFile() (err error) {
	if c.filePath == "" {
		return
	}

	dir, file := path.Split(c.filePath)
	if arr := strings.Split(file, "."); len(arr) == 2 {
		viper.SetConfigName(arr[0])
		viper.SetConfigType(arr[1])
	}

	viper.AddConfigPath(dir)

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	parseConfig := func() (err error) {
		for k, v := range c.configs {
			if k == "" {
				if err = viper.Unmarshal(v); err != nil {
					return
				}
			}
			if err = viper.UnmarshalKey(k, v); err != nil {
				return
			}
		}

		return
	}

	if err = parseConfig(); err != nil {
		return
	}

	//watch 热更新
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err = parseConfig(); err != nil {
			return
		}
	})
	viper.WatchConfig()

	return
}

//从环境变量载入配置
func (c *Config) LoadFromEnv() (err error) {
	el := EnvironmentLoader{
		CamelCase: c.envCamelCase,
	}
	for k, v := range c.configs {
		el.Prefix = strings.ToUpper(c.envPrefix)
		if k != "" {
			el.Prefix += fmt.Sprintf("_%s", strings.ToUpper(k))
		}

		if err = el.Load(v); err != nil {
			return
		}
	}
	return
}

//载入默认配置
func (c *Config) LoadFromDefault() (err error) {
	d := NewDefault()
	for _, v := range c.configs {
		if err = d.Load(v); err != nil {
			return
		}
	}
	return
}
