package redis

import (
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/golang/mock/gomock"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"

	"git.internal.yunify.com/benchmark/benchmark/mock/mock_redis"
)

//mockgen -destination=mock_redis.go -package=redis --build_flags=--mod=mod  github.com/gomodule/redigo/redis Conn
//go test  -count 1 -cover -coverprofile=coverprofile.cov && go tool cover -html=coverprofile.cov
func TestConn_Del(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("DEL", gomock.Any()).Return("", errors.New("error")),
		mockConn.EXPECT().Do("DEL", gomock.Any()).Return("", nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		keys []string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"error", fields{mockConn, 1}, args{[]string{"", ""}}, true},
		{"success", fields{mockConn, 1}, args{[]string{"", ""}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.Del(tt.args.keys...)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_Exists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), errors.New("error")),
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), nil),
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(1), nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{"redis error", fields{mockConn, 1}, args{""}, false, true},
		{"not exists", fields{mockConn, 1}, args{""}, false, false},
		{"exists", fields{mockConn, 1}, args{""}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			got, err := c.Exists(tt.args.key)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("GET", gomock.Any()).Return("", errors.New("error")),
		mockConn.EXPECT().Do("GET", gomock.Any()).Return(`"1"`, nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).Return("1", nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).Return(`{"Id":1,"Name":"name"}`, nil),
	)

	type user struct {
		Id   int
		Name string
	}

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key   string
		value interface{}
	}

	var (
		s string
		i int
		u user
	)

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"redigo: nil returned", fields{mockConn, 1}, args{"", &s}, true},
		{"string success", fields{mockConn, 1}, args{"", &s}, false},
		{"int success", fields{mockConn, 1}, args{"", &i}, false},
		{"struct success", fields{mockConn, 1}, args{"", &u}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.Get(tt.args.key, tt.args.value)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_HDel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("HDEL", gomock.Any()).Return("", errors.New("error")),
		mockConn.EXPECT().Do("HDEL", gomock.Any()).Return("0", nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key   string
		filed string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"redigo: nil returned", fields{mockConn, 1}, args{"", ""}, true},
		{"success", fields{mockConn, 1}, args{"", ""}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.HDel(tt.args.key, tt.args.filed)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_HGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("HGET", gomock.Any()).Return("", errors.New("error")),
		mockConn.EXPECT().Do("HGET", gomock.Any()).Return("0", nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key   string
		filed string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantValue string
		wantErr   bool
	}{
		{"redigo: nil returned", fields{mockConn, 1}, args{"", ""}, "", true},
		{"success", fields{mockConn, 1}, args{"", ""}, "0", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			gotValue, err := c.HGet(tt.args.key, tt.args.filed)
			assert.Equal(t, tt.wantValue, gotValue)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_HSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("HSET", gomock.Any()).Return("", errors.New("error")),
		mockConn.EXPECT().Do("HSET", gomock.Any()).Return("0", nil),
	)

	// 3. 给 hsetScript.Do 函数打桩
	outputs := []gomonkey.OutputCell{
		{
			Values: gomonkey.Params{nil, errors.New("error")},
		},
		{
			Values: gomonkey.Params{nil, nil},
		},
	}

	//对结构体及其方法进行打桩
	patches := gomonkey.ApplyMethodSeq(reflect.TypeOf(hsetScript), "Do", outputs)
	// 执行完毕之后释放桩序列
	defer patches.Reset()

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key           string
		filed         string
		value         string
		expireSeconds []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"with expire error", fields{mockConn, 1}, args{"", "", "", []int{1}}, true},
		{"with expire success", fields{mockConn, 1}, args{"", "", "", []int{1}}, false},
		{"with not expire error", fields{mockConn, 1}, args{"", "", "", []int{0}}, true},
		{"with not expire success", fields{mockConn, 1}, args{"", "", "", []int{0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.HSet(tt.args.key, tt.args.filed, tt.args.value, tt.args.expireSeconds...)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestConn_INCR(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("INCR", gomock.Any()).Return(int64(0), errors.New("error")),
		mockConn.EXPECT().Do("INCR", gomock.Any()).Return(int64(1), nil),
		mockConn.EXPECT().Do("EXPIRE", gomock.Any()).Return(nil, nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key     string
		expires []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantNum int64
		wantErr bool
	}{
		{"redis error", fields{mockConn, 1}, args{"", []int{1}}, 0, true},
		{"success", fields{mockConn, 1}, args{"", []int{1}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			gotNum, err := c.INCR(tt.args.key, tt.args.expires...)
			assert.Equal(t, tt.wantNum, gotNum)
			assert.Equal(t, tt.wantErr, err != nil)

		})
	}
}

func TestConn_Remember(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		//用例1
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), errors.New("error")),
		mockConn.EXPECT().Do("SET", gomock.Any()).Return("", nil),

		//用例2
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), errors.New("error")),

		//用例3
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), nil),
		mockConn.EXPECT().Do("SET", gomock.Any()).Return("", nil),

		//用例4
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(1), nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).Return(`"aaa"`, nil),

		//用例5
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(1), nil),
		mockConn.EXPECT().Do("GET", gomock.Any()).Return(`"aaa"`, errors.New("error")),
		mockConn.EXPECT().Do("SET", gomock.Any()).Return("", nil),

		//用例7
		mockConn.EXPECT().Do("EXISTS", gomock.Any()).Return(int64(0), nil),
	)

	var value string

	getValue := func() (interface{}, error) {
		return "aaa", nil
	}
	getValueErr := func() (interface{}, error) {
		return "aaa", errors.New("error")
	}

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key           string
		value         interface{}
		getValue      func() (interface{}, error)
		ignoreRdsErr  bool
		expireSeconds []int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"exist error ignore", fields{mockConn, 1}, args{"", &value, getValue, true, []int{10}}, false},
		{"exist error not ignore", fields{mockConn, 1}, args{"", &value, getValue, false, []int{10}}, true},
		{"exist no error and empty", fields{mockConn, 1}, args{"", &value, getValue, true, []int{10}}, false},
		{"exist no error and not emtpy", fields{mockConn, 1}, args{"", &value, getValue, true, []int{10}}, false},
		{"exist no error and not emtpy", fields{mockConn, 1}, args{"", &value, getValue, true, []int{10}}, false},
		{"getvalue error", fields{mockConn, 1}, args{"", &value, getValueErr, true, []int{}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.Remember(tt.args.key, tt.args.value, tt.args.getValue, tt.args.ignoreRdsErr, tt.args.expireSeconds...)
			assert.Equal(t, tt.wantErr, err != nil)
			if err == nil {
				assert.Equal(t, value, "aaa")
				value = ""
			}
		})
	}

}

func TestConn_Set(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Do("SET", gomock.Any()).Return(int64(0), errors.New("error")),
		mockConn.EXPECT().Do("SET", gomock.Any()).Return(int64(1), nil),
		mockConn.EXPECT().Do("SET", gomock.Any()).Return(nil, nil),
	)

	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}
	type args struct {
		key           string
		value         interface{}
		expireSeconds []int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"marshal error", fields{mockConn, 1}, args{"", make(chan bool), []int{1}}, true},
		{"redis error", fields{mockConn, 1}, args{"", "1", []int{1}}, true},
		{"success with expire", fields{mockConn, 1}, args{"", "1", []int{1}}, false},
		{"success with not expire", fields{mockConn, 1}, args{"", "1", []int{0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.Set(tt.args.key, tt.args.value, tt.args.expireSeconds...)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestNewConn(t *testing.T) {
	assert.NotNil(t, NewConn(nil))
}

func TestConn_Close(t *testing.T) {
	type fields struct {
		conn                 redis.Conn
		defaultExpireSeconds int
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 1. 生成符合 redis.Conn 接口的 mockConn
	mockConn := mock_redis.NewMockConn(ctrl)
	// 2. 给接口打桩序列
	gomock.InOrder(
		mockConn.EXPECT().Close().Return(errors.New("error")),
		mockConn.EXPECT().Close().Return(nil),
	)

	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"error", fields{mockConn, 1}, true},
		{"success", fields{mockConn, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conn{
				conn:                 tt.fields.conn,
				defaultExpireSeconds: tt.fields.defaultExpireSeconds,
			}
			err := c.Close()
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
