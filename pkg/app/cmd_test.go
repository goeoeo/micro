package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCmd(t *testing.T) {
	c := NewCmd(func(c *Cmd) (err error) {
		return nil
	})

	c.Register("path", "p", "", "配置文件路径")
	assert.Equal(t, c.Get("path", ""), "")
	assert.Equal(t, c.ConfigPath(""), "")
	assert.Equal(t, c.ConfigPath("aaa"), "aaa")
	assert.Nil(t, c.Run())

}
