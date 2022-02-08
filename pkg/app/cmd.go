package app

import (
	"github.com/spf13/cobra"
)

type Cmd struct {
	params     map[string]*string
	configPath string
	cmd        *cobra.Command
}

func NewCmd(run func(c *Cmd) (err error)) *Cmd {
	c := &Cmd{
		params: make(map[string]*string),
		cmd: &cobra.Command{
			Use:   "",
			Short: "default",
		},
	}

	c.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return run(c)
	}

	c.cmd.Flags().StringVarP(&c.configPath, "configPath", "c", "", "配置文件路径")
	return c
}

//注册参数
func (c *Cmd) Register(name string, shorthand, value, usage string) *Cmd {
	var tmp string
	c.params[name] = &tmp
	c.cmd.Flags().StringVarP(c.params[name], name, shorthand, value, usage)
	return c
}

//获取参数
func (c *Cmd) Get(key string, value string) string {
	if tmp, ok := c.params[key]; ok {
		return *tmp
	}

	return value
}

func (c *Cmd) Run() error {
	return c.cmd.Execute()
}

//ConfigPath 配置文件路径
func (c *Cmd) ConfigPath(path string) string {
	if c != nil && c.configPath != "" {
		return c.configPath
	}

	return path
}
