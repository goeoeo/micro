package redis

import (
	"errors"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"

	"git.internal.yunify.com/benchmark/benchmark/mock/mock_redis"
)

func TestNewRedisCli(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Close().Return(nil),
	)

	// 3. 给 redis.Dail 函数打桩
	outputs := []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{nil, errors.New("error")},
		},
		{
			Values: gomonkey.Params{mockConn, nil},
		},
		{
			Values: gomonkey.Params{mockConn, nil},
		},
	}

	patches := gomonkey.ApplyFuncSeq(redis.Dial, outputs)
	// 执行完毕之后释放桩序列
	defer patches.Reset()

	cli := NewRedisCli("redis")

	assert.NotNil(t, cli.Init())
	assert.Nil(t, cli.Init())
	assert.NotNil(t, cli.Cache())

	assert.NotNil(t, cli.NewMutex("aaa"))

	name, config := cli.Config()
	assert.Equal(t, name, "redis")
	assert.NotNil(t, config)

}
