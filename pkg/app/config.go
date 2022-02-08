package app

import (
	"fmt"

	"git.internal.yunify.com/benchmark/benchmark/pkg/util/iputil"
	"git.internal.yunify.com/benchmark/benchmark/pkg/util/netutil"
)

//当前项目环境
const (
	ModeDev     = "dev"
	ModeStaging = "staging"
	ModeRelease = "release"
)

type Config struct {
	Project string `default:"Project"` //项目
	Service string `default:"Service"` //服务
	Version string `default:"v0.0.1"`  //版本
	Mode    string `default:"dev"`     //环境: dev=开发环境 staging=测试环境 release=生产环境
	Host    string `default:""`
	Port    string `default:"8080"`
}

func (c *Config) Endpoint() string {
	if c.Host == "" {
		ip, _ := iputil.ExternalIP()

		//获取本地网卡ip
		c.Host = fmt.Sprintf("%s", ip)
	}
	return netutil.HostPort(c.Host, c.Port)
}

func (c *Config) IsDev() bool {
	return c.Mode == ModeDev
}

func (c *Config) IsStaging() bool {
	return c.Mode == ModeStaging
}

func (c *Config) IsModeRelease() bool {
	return c.Mode == ModeRelease
}
