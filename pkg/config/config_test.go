package config

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

type (
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
		Name struct{
			A string `default:"aaa"`
		}
	}
)

var configTest="./config_test.toml"

func TestConfig_loadFromFile(t *testing.T) {
	var(
		configTestContent []byte
		configTestString string
		err error
		mysql MysqlConfig
	)
	err=NewConfig(configTest).Register("mysql",&mysql).loadFromFile()
	assert.Nil(t, err)

	assert.Equal(t, mysql.Host,"192.168.49.2")
	assert.Equal(t, mysql.Name.A,"bbb")

	change:= func(a,b string) {
		configTestContent,err=ioutil.ReadFile(configTest)
		assert.Nil(t, err)

		configTestString=strings.Replace(string(configTestContent),a,b,-1)
		err=ioutil.WriteFile(configTest,[]byte(configTestString),os.ModePerm)

		assert.Nil(t,err)
		//更新是异步的这里停顿一段时间
		time.Sleep(10*time.Millisecond)
	}


	//热更新
	change(`host = "192.168.49.2"`,`host = "192.168.49.3"`)
	assert.Equal(t, mysql.Host,"192.168.49.3")

	change(`host = "192.168.49.3"`,`host = "192.168.49.2"`)
	assert.Equal(t, mysql.Host,"192.168.49.2")
}

func TestConfig_loadFromEnv(t *testing.T)  {
	var(
		err error
		mysql MysqlConfig
	)
	_=os.Setenv("MYSQL_HOST","192.168.49.2")
	_=os.Setenv("MYSQL_NAME_A","bbb")

	err=NewConfig(configTest).Register("mysql",&mysql).loadFromEnv()
	assert.Nil(t, err)

	assert.Equal(t, mysql.Host,"192.168.49.2")
	assert.Equal(t, mysql.Name.A,"bbb")

}

func TestConfig_loadFromDefault(t *testing.T)  {
	var(
		err error
		mysql MysqlConfig
	)

	err=NewConfig(configTest).Register("mysql",&mysql).loadFromDefault()
	assert.Nil(t, err)

	assert.Equal(t, mysql.Host,"localhost")
	assert.Equal(t, mysql.Name.A,"aaa")

}