package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

type Conn struct {
	conn                 redis.Conn
	defaultExpireSeconds int //默认过期时间 单位秒
}

func NewConn(conn redis.Conn) *Conn {
	return &Conn{
		conn:                 conn,
		defaultExpireSeconds: 24 * 3600,
	}
}

func (c *Conn) Remember(key string, value interface{}, getValue func() (interface{}, error), ignoreRdsErr bool, expireSeconds ...int) (err error) {
	getValueFromSource := func() (err error) {
		var getValueRes interface{}
		if getValueRes, err = getValue(); err != nil {
			return
		}

		//不通对象地址，marshal转换
		if fmt.Sprintf("%p", value) != fmt.Sprintf("%p", getValueRes) {
			var data []byte
			if data, err = jsoniter.Marshal(getValueRes); err != nil {
				return
			}
			if err = jsoniter.Unmarshal(data, value); err != nil {
				return
			}
		}

		//异步存入redis
		if err = c.Set(key, getValueRes, expireSeconds...); err != nil {
			return
		}

		return
	}

	var exist bool
	if exist, err = c.Exists(key); err != nil {
		if ignoreRdsErr {
			err = nil
			return getValueFromSource()
		}

		return
	}

	if !exist {
		return getValueFromSource()
	}

	if err = c.Get(key, value); err != nil {
		if ignoreRdsErr {
			err = nil
			return getValueFromSource()
		}
		return
	}

	return
}

func (c *Conn) Get(key string, value interface{}) (err error) {
	var data string
	if data, err = redis.String(c.conn.Do("GET", key)); err != nil {
		return
	}

	return jsoniter.Unmarshal([]byte(data), value)
}

func (c *Conn) Set(key string, value interface{}, expireSeconds ...int) (err error) {
	var (
		data []byte
	)

	expire := c.getExpire(expireSeconds...)

	//序列化数据
	if data, err = jsoniter.Marshal(value); err != nil {
		return
	}

	if expire > 0 {
		if _, err = c.conn.Do("SET", key, string(data), "EX", fmt.Sprintf("%d", expire)); err != nil {
			return err
		}
	} else {
		if _, err = c.conn.Do("SET", key, string(data)); err != nil {
			return err
		}
	}

	return
}

func (c *Conn) Del(keys ...string) (err error) {
	if len(keys) == 0 {
		return
	}

	var keysi []interface{}
	for _, v := range keys {
		keysi = append(keysi, v)
	}

	if _, err = c.conn.Do("DEL", keysi...); err != nil {
		return
	}

	return
}

func (c *Conn) Exists(key string) (bool, error) {
	return redis.Bool(c.conn.Do("EXISTS", key))
}

func (c *Conn) INCR(key string, expires ...int) (num int64, err error) {
	if num, err = redis.Int64(c.conn.Do("INCR", key)); err != nil {
		return
	}

	expire := c.getExpire(expires...)
	if num == 1 && expire > 0 {
		_, err = c.conn.Do("EXPIRE", key, expire)
	}

	return
}

func (c *Conn) HSet(key, filed, value string, expireSeconds ...int) (err error) {

	expire := c.getExpire(expireSeconds...)

	//需要设置过期时间
	if expire > 0 {
		if _, err = hsetScript.Do(c.conn, key, filed, value, expire); err != nil {
			return err
		}
		return
	}

	if _, err = c.conn.Do("HSET", key, filed, value); err != nil {

		return err
	}

	return
}

func (c *Conn) HGet(key, filed string) (value string, err error) {
	value, err = redis.String(c.conn.Do("HGET", key, filed))
	return
}

func (c *Conn) HDel(key, filed string) (err error) {
	_, err = c.conn.Do("HDEL", key, filed)
	return
}

func (c *Conn) Close() error {
	return c.conn.Close()
}

//获取过期时间
func (r *Conn) getExpire(expire ...int) int {
	if len(expire) > 0 {
		return expire[0]
	}

	return r.defaultExpireSeconds
}

func (c *Conn) ZSet(key, value string, score float64, expireSeconds ...int) (err error) {

	expire := c.getExpire(expireSeconds...)

	if expire > 0 {
		if _, err = zaddScript.Do(c.conn, key, score, value, expire); err != nil {
			return
		}

		return
	}

	if _, err = c.conn.Do("ZADD", key, score, value); err != nil {
		return
	}

	return
}

func (c *Conn) ZGet(key string, limit int, asc ...bool) (res []string, err error) {
	var (
		commandName = "ZREVRANGE" //默认降序
	)

	if len(asc) > 0 && asc[0] {
		commandName = "ZRANGE"
	}

	if res, err = redis.Strings(c.conn.Do(commandName, key, 0, limit)); err != nil {
		return
	}

	return
}
