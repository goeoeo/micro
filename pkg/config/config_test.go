package config

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type (
	ConfigT struct {
		A     string
		Mysql MysqlConfig
	}
	MysqlConfig struct {
		Host            string `default:"localhost"`
		Port            string `default:"3306"`
		User            string `default:"root"`
		Password        string `default:"123456"`
		Database        string `default:""`
		MaxIdleConns    int    `default:"100"`
		MaxOpenConns    int    `default:"100"`
		ConnMaxLifetime int    `default:"10"` //time.Second
		Disable         bool   `default:"false"`
		Name            struct {
			A string `default:"aaa"`
		}
	}
)

var (
	filePath   = "./config_test.toml"
	configTest = WithFilePath(filePath)
)

func TestConfig_loadFromFile(t *testing.T) {
	var (
		configTestContent []byte
		configTestString  string
		err               error
		mysql             MysqlConfig
		c                 ConfigT
	)

	//改变文件
	change := func(a, b string) {
		configTestContent, err = ioutil.ReadFile(filePath)
		assert.Nil(t, err)

		configTestString = strings.Replace(string(configTestContent), a, b, -1)
		err = ioutil.WriteFile(filePath, []byte(configTestString), os.ModePerm)

		assert.Nil(t, err)
		//更新是异步的这里停顿一段时间
		time.Sleep(100 * time.Millisecond)
	}

	t.Run("section不为空", func(t *testing.T) {
		err = NewConfig(configTest).Register("mysql", &mysql).LoadFromFile()
		assert.Nil(t, err)

		assert.Equal(t, mysql.Host, "192.168.49.2")
		assert.Equal(t, mysql.Name.A, "bbb")

		//热更新
		change(`host = "192.168.49.2"`, `host = "192.168.49.3"`)
		assert.Equal(t, mysql.Host, "192.168.49.3")

		change(`host = "192.168.49.3"`, `host = "192.168.49.2"`)
		assert.Equal(t, mysql.Host, "192.168.49.2")
	})

	t.Run("section为空", func(t *testing.T) {
		err = NewConfig(configTest).Register("", &c).LoadFromFile()
		assert.Nil(t, err)

		assert.Equal(t, c.Mysql.Host, "192.168.49.2")
		assert.Equal(t, c.Mysql.Name.A, "bbb")

		//热更新
		change(`host = "192.168.49.2"`, `host = "192.168.49.3"`)
		assert.Equal(t, c.Mysql.Host, "192.168.49.3")

		change(`host = "192.168.49.3"`, `host = "192.168.49.2"`)
		assert.Equal(t, c.Mysql.Host, "192.168.49.2")

	})

}

func TestConfig_loadFromEnv(t *testing.T) {
	var (
		err   error
		mysql MysqlConfig
		c     ConfigT

		prefix = WithEnvPrefix("NB")
	)
	_ = os.Setenv("NB_MYSQL_HOST", "192.168.49.2")
	_ = os.Setenv("NB_MYSQL_NAME_A", "bbb")

	t.Run("section不为空", func(t *testing.T) {
		err = NewConfig(prefix).Register("mysql", &mysql).LoadFromEnv()
		assert.Nil(t, err)

		assert.Equal(t, mysql.Host, "192.168.49.2")
		assert.Equal(t, mysql.Name.A, "bbb")
	})

	t.Run("section为空", func(t *testing.T) {
		err = NewConfig(prefix).Register("", &c).LoadFromEnv()
		assert.Nil(t, err)

		assert.Equal(t, c.Mysql.Host, "192.168.49.2")
		assert.Equal(t, c.Mysql.Name.A, "bbb")
	})

}

func TestConfig_loadFromDefault(t *testing.T) {
	var (
		err   error
		mysql MysqlConfig
		c     ConfigT
	)

	t.Run("section不为空", func(t *testing.T) {
		err = NewConfig().Register("mysql", &mysql).LoadFromDefault()
		assert.Nil(t, err)

		assert.Equal(t, mysql.Host, "localhost")
		assert.Equal(t, mysql.Name.A, "aaa")
	})

	t.Run("section为空", func(t *testing.T) {
		err = NewConfig().Register("", &c).LoadFromDefault()
		assert.Nil(t, err)

		assert.Equal(t, c.Mysql.Host, "localhost")
		assert.Equal(t, c.Mysql.Name.A, "aaa")
	})

}
